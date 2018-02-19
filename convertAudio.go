package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func convertAudio(audioPath string, duration int) (string, error) {
	//ffmpeg -stream_loop -1 -i in.mp3 -c copy -t 600 ./out.mp3  # 600秒間ループ
	outAudioPath, _ := filepath.Abs(filepath.Join(`./`, strconv.FormatInt(time.Now().UnixNano(), 10)+".mp3"))

	args := []string{
		"-stream_loop", "-1", "-i", audioPath, "-acodec", "copy", "-t", strconv.Itoa(duration), outAudioPath,
	}

	out, err := exec.Command("ffmpeg", args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ConvertAudio error: %s, %s", err, out)
	}

	return outAudioPath, nil
}
