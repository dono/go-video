package main

import (
	"log"
	"testing"
)

func TestTextToImage(t *testing.T) {
	_, err := TextToImage(`./assets/sample.txt`, `./out.png`, `/Library/Fonts/Ricty-Regular.ttf`, 25, 800)
	if err != nil {
		log.Fatal(err)
	}
}
