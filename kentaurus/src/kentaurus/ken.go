package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewKentaurus() (kens []Ken) {

	kens = []Ken{
		{
			img:  kentaurus2Image,
			d:    4,
			size: screenWidth / 12,
			y:    380,
		},
		{
			img:  kentaurus3Image,
			d:    3,
			size: screenWidth / 10,
			y:    400,
		},
		{
			img:  kentaurus1Image,
			d:    1,
			size: screenWidth / 8,
			y:    420,
		},
	}
	return
}

type Ken struct {
	img  *ebiten.Image
	size int
	d    int
	x    int
	y    int
}

func (k *Ken) Update(ctr int) {
	// 背景の中心に左右にぶらぶらさせる
	rad := float64(ctr*2*k.d) * math.Pi / MaxCtr
	k.x = screenWidth/2 + int(float64(k.size)*math.Cos(rad))
	/*
		nextX := k.x + k.dx
		nextDX := -(k.x - centerX) / 2

		k.x = nextX
		k.dx = nextDX
	*/
}

func (k *Ken) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := k.img.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Translate(float64(k.x), float64(k.y))
	screen.DrawImage(k.img, op)
}
