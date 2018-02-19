package main

import (
	"fmt"
	"log"
	"testing"
)

func TestConvertAudio(t *testing.T) {
	path, err := convertAudio("./assets/audio.mp3", 200)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}
