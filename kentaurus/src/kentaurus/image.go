package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"images"
)

var (
	kfcImage        *ebiten.Image // KFCの背景画像
	kentaurus1Image *ebiten.Image // ケンタウルスの画像1
	kentaurus2Image *ebiten.Image // ケンタウルスの画像2
	kentaurus3Image *ebiten.Image // ケンタウルスの画像3
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.KfcPng))
	if err != nil {
		log.Fatal(err)
	}
	kfcImage = ebiten.NewImageFromImage(img)

	if kfcImage == nil {
		log.Fatal("kfcImage is empty")
	}

	img, _, err = image.Decode(bytes.NewReader(images.Kentaurus1Png))
	if err != nil {
		log.Fatal(err)
	}
	kentaurus1Image = ebiten.NewImageFromImage(img)

	if kentaurus1Image == nil {
		log.Fatal("kentaurus1Image is empty")
	}

	img, _, err = image.Decode(bytes.NewReader(images.Kentaurus2Png))
	if err != nil {
		log.Fatal(err)
	}
	kentaurus2Image = ebiten.NewImageFromImage(img)

	if kentaurus2Image == nil {
		log.Fatal("kentaurus2Image is empty")
	}

	img, _, err = image.Decode(bytes.NewReader(images.Kentaurus3Png))
	if err != nil {
		log.Fatal(err)
	}
	kentaurus3Image = ebiten.NewImageFromImage(img)

	if kentaurus3Image == nil {
		log.Fatal("kentaurus3Image is empty")
	}
}
