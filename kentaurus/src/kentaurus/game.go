package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	MaxCtr = 360
)

func NewGame() (g *Game) {
	g = &Game{
		kens: NewKentaurus(),
	}
	return
}

type Game struct {
	ctr  int // 0 ≦ ctr < MaxCtr
	kens []Ken
}

func (g *Game) Update() (e error) {

	for i := 0; i < len(g.kens); i++ {
		g.kens[i].Update(g.ctr)
	}

	g.ctr = (g.ctr + 1) % MaxCtr

	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 背景色
	//ebitenutil.DrawRect(screen, 0, 0, screenWidth, screenHeight, skyblueColor)

	// KFC背景画像を配置
	screen.DrawImage(kfcImage, &ebiten.DrawImageOptions{})

	for i := 0; i < len(g.kens); i++ {
		g.kens[i].Draw(screen)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("x = %d, dx = %d", g.kens[i].x, g.kens[i].d))
	}

	/*
		op := &ebiten.DrawImageOptions{}
		w, h := jesusIcoImage.Size()
		//fmt.Println("[debug]", "w =", w, "h =", h)

		// スクリーンの中心に移動
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)

		// スクリーンの中心からの座標を決定
		nextX, nextY := g.nextPosition()
		op.GeoM.Translate(nextX, nextY)

		// 善次を配置
		screen.DrawImage(jesusIcoImage, op)

		// マウスがクリックされてるとき
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			op2 := &ebiten.DrawImageOptions{}
			w2, h2 := paraisoImage.Size()
			// スクリーンの中心に移動
			op2.GeoM.Translate(-float64(w2)/2, -float64(h2)/2)
			op2.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)
			// ぱらいそのコマを配置
			screen.DrawImage(paraisoImage, op2)
		}

	*/

	//fmt.Println("[debug]", "nextX =", nextX, "nextY =", nextY) // for debug
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

/*
func (g *Game) nextPosition() (nextX, nextY float64) {
	t := float64(g.ctr) / 50
	nextX = math.Cos(t) * 100
	nextY = math.Sin(t*2) * 50
	return

	//	x = cos(t), y = sin(2t) で「８の字」軌道になる
}
*/
