package main

import (
	"fmt"
	"image"
	"os/exec"
	"strconv"

	"github.com/disintegration/imaging"
)

func TextToImage(textPath string, outTextImgPath string, fontPath string, pointSize int, tImgW int) (image.Image, error) {
	// convert -size 800x100000 xc:none -font /Library/Fonts/Ricty-Regular.ttf -pointsize 30 -fill white -stroke black -strokewidth 7 -annotate +0+100 @sample.txt -strokewidth 0 -stroke white -annotate +0+100 @sample.txt -trim +repage new.png

	var tImg image.Image
	args := []string{
		"-size", strconv.Itoa(tImgW) + "x100000",
		"xc:none",
		"-font", fontPath,
		"-pointsize", strconv.Itoa(pointSize),
		"-fill", "white",
		"-stroke", "black",
		"-strokewidth", "7",
		"-annotate", "+0+100", "@" + textPath,
		"-strokewidth", "0",
		"-stroke", "white",
		"-annotate", "+0+100", "@" + textPath,
		"-trim", "+repage",
		outTextImgPath,
	}

	out, err := exec.Command("convert", args...).CombinedOutput()
	if err != nil {
		return tImg, fmt.Errorf(string(out))
	}

	return imaging.Open(outTextImgPath)
}
