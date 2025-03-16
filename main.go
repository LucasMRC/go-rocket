package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	rocket_main := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "rocket",
		Width:         130,
		Height:        48,
		MinWidth:      0,
		MinHeight:     0,
		AlwaysOnTop:   true,
		Frameless:     true,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: options.NewRGBA(0, 0, 0, 0),
		Linux: &linux.Options{
			WindowIsTranslucent: true,
			ProgramName:         "Rocket",
		},
		OnStartup: rocket_main.startup,
		Bind: []any{
			rocket_main,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
