package main

import "os/exec"

// 動画ファイルと音声ファイルを合成する
func combineVideoAndAudio(videoPath, audioPath, outVideoPath string) error {
	args := []string{
		"-i", videoPath, "-i", audioPath, "-c", "copy", outVideoPath,
	}

	err := exec.Command("ffmpeg", args...).Run()
	if err != nil {
		return err
	}

	return nil
}
