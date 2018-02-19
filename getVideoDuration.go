package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

// 動画の長さを秒単位で返す ミリ秒は切り捨て
func getVideoDuration(videoPath string) (int, error) {
	duration := -1

	args := []string{
		"-i", videoPath,
	}

	out, err := exec.Command("ffprobe", args...).CombinedOutput()
	if err != nil {
		return 1, fmt.Errorf("GetVideoDuration error: %s", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Duration") {
			str := strings.Split(line, " ")[3] // HH:mm:ss.ff,

			var hour, min, sec int
			_, err := fmt.Sscanf(str, "%d:%d:%d", &hour, &min, &sec)
			if err != nil {
				return -1, err
			}

			duration = hour*60*60 + min*60 + sec
		}
	}

	return duration, nil
}
