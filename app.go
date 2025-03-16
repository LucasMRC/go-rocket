package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}
type ProcessStatus int
type AppResponse[T any] struct {
	Status   ProcessStatus `json:"status"`
	Response T             `json:"response"`
}

const (
	ProcessSuccess ProcessStatus = iota
	ProcessRetry
	ProcessCanceled
	ProcessError
	RecordingStartedStdout = "Input #0, x11grab, from"
)

var (
	ctx            context.Context
	ffmpeg_stdin   io.WriteCloser
	file_temp_name = os.TempDir() + "/output.mp4"
	cmd            *exec.Cmd
	done           = make(chan bool)
)

func NewApp() *App {
	fmt.Println(cmd)

	return &App{}
}

func (a *App) startup(actx context.Context) {
	a.ctx = actx
	ctx = actx
}

func (a *App) GetScreenObjects() AppResponse[[]runtime.Screen] {
	screens, err := runtime.ScreenGetAll(ctx)
	if err != nil {
		fmt.Println("ERRORRR", err.Error())
		return AppResponse[[]runtime.Screen]{
			Response: []runtime.Screen{},
			Status:   ProcessError,
		}
	}
	return AppResponse[[]runtime.Screen]{
		Response: screens,
		Status:   ProcessSuccess,
	}
}

func (a *App) StartRecording() AppResponse[any] {
	fmt.Println("Start recording!")
	cmd = exec.CommandContext(context.TODO(), "ffmpeg", "-select_region", "1", "-framerate", "25", "-f", "x11grab", "-i", ":0.0", "-y", "/tmp/output.mp4")

	stdout, _ := cmd.StdoutPipe()
	ffmpeg_stdin, _ = cmd.StdinPipe()
	cmd.Stderr = cmd.Stdout
	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, RecordingStartedStdout) {
				SendRecordingSignal()
			}
		}
	}()
	err := cmd.Start()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
		<-done
		return AppResponse[any]{
			Response: nil,
			Status:   ProcessError,
		}
	}
	return AppResponse[any]{
		Status:   ProcessSuccess,
		Response: nil,
	}
}

func (a *App) StopRecording() AppResponse[any] {
	fmt.Println("Stop recording!")
	defer ffmpeg_stdin.Close()
	_, err := io.WriteString(ffmpeg_stdin, "q")
	if err != nil {
		fmt.Println("ERRRORR cancel", err.Error())
		return AppResponse[any]{
			Status:   ProcessError,
			Response: nil,
		}
	}
	path, err := runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{
		DefaultFilename: "output.mp4",
		Title:           "Save recording",
	})
	if err != nil {
		fmt.Println("ERRRORR save", err.Error())
		return AppResponse[any]{
			Status:   ProcessError,
			Response: nil,
		}
	}
	if path != "" {
		mv_cmd := exec.Command("mv", file_temp_name, path)
		err = mv_cmd.Run()
		if err != nil {
			fmt.Println("ERRRORR mv_cmd", err)
			return AppResponse[any]{
				Status:   ProcessError,
				Response: nil,
			}
		}
		mv_cmd.Wait()
	}
	return AppResponse[any]{
		Status:   ProcessSuccess,
		Response: nil,
	}
}

func SendRecordingSignal() {
	runtime.EventsEmit(ctx, "recording-started", nil)
}
