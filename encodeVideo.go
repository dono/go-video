package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

// 連番画像と音楽を元に動画を生成
func encodeVideo(imagesDirPath string) (string, error) {
	outVideoPath, _ := filepath.Abs(filepath.Join(`./`, strconv.FormatInt(time.Now().UnixNano(), 10)+".mp4"))
	videoFPS := 30

	args := []string{
		"-framerate", strconv.Itoa(imgFPS),
		"-i", filepath.Join(imagesDirPath, "%d.png"),
		"-vcodec", "libx264",
		"-pix_fmt", "yuv420p",
		"-r", strconv.Itoa(videoFPS),
		"-vf", "fade=t=in:st=0:d=1",
		"-loglevel", "info",
		outVideoPath,
	}

	err := exec.Command("ffmpeg", args...).Run()
	if err != nil {
		return "", fmt.Errorf("EncodeVideo error: %s", err)
	}

	return outVideoPath, nil
}
