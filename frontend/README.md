	// WebCam: ffmpeg -f v4l2 -framerate 60 -video_size 1280x720 -input_format mjpeg -i /dev/video1 -preset faster -pix_fmt yuv420p out.mkv
	// ffmpeg -select_region 1 -framerate 25 -f x11grab -i :0.0  output.mp4
	// Mic: ffmpeg -f alsa -i hw:1 -t 30 out.wav
