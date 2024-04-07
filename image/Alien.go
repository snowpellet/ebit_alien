package image

import (
	"bytes"
	config_test "ebit/config"
	"ebit/resourse"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// 绘制外星人
type Alien struct {
	Image       *ebiten.Image `json:"img"`
	Width       int           `json:"width"`
	Height      int           `json:"height"`
	X           float64       `json:"X"`
	Y           float64       `json:"Y"`
	SpendFactor int           `json:"spendFactor"`
}

func NewAlien(cfg *config_test.Config) *Alien {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resourse.Alien))
	if err != nil {
		log.Fatalf("img alien %v\n", err)
	}
	width, height := img.Size()
	return &Alien{
		Image:       img,
		Width:       width,
		Height:      height,
		X:           0,
		Y:           0,
		SpendFactor: cfg.SpendFactor,
	}
}
func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	//type Drawimageoptions struct
	// GeoM是一个要绘制的几何矩阵。
	//默认(0)值是identity，它在(0,0)处绘制图像。
	// Translate通过(tx, ty)平移矩阵。

	op.GeoM.Translate(alien.X, alien.Y)
	screen.DrawImage(alien.Image, op)
}
