package main

import (
	"context"
	"image"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"golang.org/x/sync/errgroup"
)

// テキスト画像と背景画像を合成する 戻り値は保存先ディレクトリパス
func generateImages(tImg, bg image.Image) (string, error) {
	ppf := vel / imgFPS // 1フレームあたりに移動させるピクセル数 (px/frame)
	bgW := vidW         // 背景画像の幅   (px)
	bgH := vidH         // 背景画像の高さ (px)
	tImgH := tImg.Bounds().Dy()
	cntFrames := (tImgH + vidH) / ppf

	// 背景画像のリサイズ
	bgBounds := bg.Bounds()
	if bgBounds.Dx() != bgW && bgBounds.Dy() != bgH {
		bg = imaging.Resize(bg, bgW, bgH, imaging.Lanczos)
		bgBounds = bg.Bounds()
	}

	// 背景画像に対するテキスト画像を合成する座標を設定
	bgMinX := bgBounds.Min.X
	bgMinY := bgBounds.Min.Y
	centerX := bgMinX + bgW/2
	x0 := centerX - tImg.Bounds().Dx()/2

	// 画像保存先のパス
	imagesDirPath, _ := filepath.Abs(filepath.Join(`./`, strconv.FormatInt(time.Now().UnixNano(), 10)))
	if err := os.Mkdir(imagesDirPath, 0755); err != nil {
		return "", err
	}

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 合成
	goLimiter := make(chan struct{}, 6) // 最大goroutine数を6に設定
	for i := 1; i < cntFrames+1; i++ {
		i := i
		eg.Go(func() error {
			goLimiter <- struct{}{}
			defer func() { <-goLimiter }()

			select {
			case <-ctx.Done():
				return nil
			default:
				y := bgMinY + bgH - ppf*i
				out := imaging.Overlay(bg, tImg, image.Pt(x0, y), 1)
				err := imaging.Save(out, filepath.Join(imagesDirPath, strconv.Itoa(i)+".png"))

				return err
			}
		})
	}
	if err := eg.Wait(); err != nil {
		cancel()
		return "", err
	}

	return imagesDirPath, nil
}
