package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Format2chText(textPath string) (string, error) {
	splitLen := 28
	outTextPath, _ := filepath.Abs(filepath.Join(`./`, strconv.FormatInt(time.Now().UnixNano(), 10)+".txt"))

	in, err := os.Open(textPath)
	if err != nil {
		return "", err
	}
	defer in.Close()

	out, err := os.OpenFile(outTextPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer out.Close()

	keywords := []string{" 2015/", " 2016/", " 2017/", " 2018/", " 2019/"}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		// 日付とIDを削除
		for _, word := range keywords {
			index := strings.Index(line, word)
			if index != -1 {
				line = line[:index]
			}
		}

		if len(line) == 0 {
			fmt.Fprintln(out, line)
			continue
		}

		// 名前の行は改行をいれない
		if !strings.Contains(line, "名無し") {
			// splitLen文字で改行を入れる
			runes := []rune(line)
			for i := 0; i < len(runes); i += splitLen {
				if i+splitLen < len(runes) {
					fmt.Fprintln(out, string(runes[i:(i+splitLen)]))
				} else {
					fmt.Fprintln(out, string(runes[i:]))
				}
			}
		} else {
			fmt.Fprintln(out, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return outTextPath, nil
}
