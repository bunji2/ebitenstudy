package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"images"
)

var (
	jesusIcoImage *ebiten.Image // 善次のアイコン画像
	paraisoImage  *ebiten.Image // ぱらいそのコマ画像
	whiteColor    color.Color
	skyblueColor  color.Color
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.JesusIcoPng))
	if err != nil {
		log.Fatal(err)
	}
	jesusIcoImage = ebiten.NewImageFromImage(img)

	img2, _, err := image.Decode(bytes.NewReader(images.ParaisoPng))
	if err != nil {
		log.Fatal(err)
	}
	paraisoImage = ebiten.NewImageFromImage(img2)

	if jesusIcoImage == nil {
		log.Fatal("jesusIcoImage is empty")
	}

	if paraisoImage == nil {
		log.Fatal("paraisoImage is empty")
	}

	skyblueColor = color.RGBA{R: 0x87, G: 0xce, B: 0xeb, A: 0xff}
}
