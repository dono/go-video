package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/disintegration/imaging"
)

const (
	imgFPS   = 24             // 生成された画像のFPS
	vidW     = 854            // videoのwidth(px)
	vidH     = 480            // height(px)
	tImgW    = 800            // テキスト画像のwidth(px)
	vel      = int(vidH / 10) // テキストの流れる速さ(px/sec)
	fontSize = 34

	// 入力
	fontFilePath = "/Library/Fonts/Ricty-Regular.ttf"
	textFilePath = "./assets/sample.txt"
	bgImagePath  = "./assets/background.jpg"
	audioPath    = "./assets/audio.mp3"

	// 出力先
	outTextImgPath = "./generated/text.png"
	outVideoPath   = "./generated/out.mp4"
)

func main() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	// テキストデータの前処理
	log.Println("Pre processing text...")
	s.Start()
	tmpTextPath, err := Format2chText(textFilePath)
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()

	// テキストファイルをを画像に変換
	log.Println("Converting text to image...")
	s.Start()
	tImg, err := TextToImage(tmpTextPath, outTextImgPath, fontFilePath, 26, tImgW)
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()

	tImgH := tImg.Bounds().Dy()
	t := tImgH / vel // テキスト画像が流れ切るのに必要な時間(s)

	fmt.Printf("サイズ:   %dx%d\n", tImgW, tImgH)
	fmt.Printf("所要時間: %dsec(%gmin)\n\n", t, float64(t)/float64(60))

	// 背景画像を読み込む
	bg, err := imaging.Open(bgImagePath)
	if err != nil {
		log.Fatal(err)
	}

	// 画像生成
	log.Println("Generating images...")
	s.Restart()
	tmpImagesDirPath, err := generateImages(tImg, bg)
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()

	// 音声無し動画生成
	log.Println("Encoding video...")
	s.Restart()
	tmpVideoPath, err := encodeVideo(tmpImagesDirPath)
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()

	// 動画の長さを取得
	duration, err := getVideoDuration(tmpVideoPath)
	if err != nil {
		log.Fatal(err)
	}

	// 音声ファイルを生成
	log.Println("Converting audio...")
	s.Restart()
	tmpAudioPath, err := convertAudio(audioPath, duration)
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()

	// 動画ファイルと音声ファイルを合成
	log.Println("Combining video and audio...")
	s.Restart()
	err = combineVideoAndAudio(tmpVideoPath, tmpAudioPath, outVideoPath)
	if err != nil {
		log.Fatal(err)
	}
	s.Stop()

	// 後処理
	defer func() {
		log.Println("Post processing...")
		s.Restart()
		if err := os.RemoveAll(tmpTextPath); err != nil {
			log.Fatal(err)
		}

		if err := os.RemoveAll(tmpImagesDirPath); err != nil {
			log.Fatal(err)
		}

		if err := os.RemoveAll(tmpVideoPath); err != nil {
			log.Fatal(err)
		}

		if err := os.RemoveAll(tmpAudioPath); err != nil {
			log.Fatal(err)
		}
		s.Stop()

		log.Println("Complete!")
	}()
}
