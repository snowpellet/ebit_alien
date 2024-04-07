package image

import (
	"bytes"
	config_test "ebit/config"
	"ebit/resourse"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	logging "github.com/sirupsen/logrus"
)

type Ship struct {
	Image  *ebiten.Image `json:"image"`
	Width  int           `json:"width"` //当前位置
	Height int           `json:"height"`
	X      int           `json:"x"`
	Y      int           `json:"y"`
}

// 返回一个图片和图片的大小
func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resourse.Ship))
	if err != nil {
		logging.Info("image:", err)
	}
	width, height := img.Size()
	ship := &Ship{
		Image:  img,
		Width:  width,
		Height: height,
		X:      (screenWidth - width) / 2, // 计算初始坐标需要屏幕尺寸，故增加屏幕宽、高两个参数，
		Y:      height,                    //
	}
	return ship
}

// 绘制自身,传入屏幕对象screen和配置，让代码更好维护：再让game.Draw调用
// 我们给Ship类型增加一个绘制自身的方法，传入屏幕对象screen和配置，让代码更好维护：
/*
	x=(W1-W2)/2界面大小减去图片大小/2
	y=H1-H2
*/
//屏幕中间显示飞船图片
func (ship *Ship) Draw(screen *ebiten.Image, cfg *config_test.Config) {
	op := &ebiten.DrawImageOptions{}
	//type Drawimageoptions struct
	// GeoM是一个要绘制的几何矩阵。
	//默认(0)值是identity，它在(0,0)处绘制图像。
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.X-ship.Width), float64(cfg.ScreenHeight-ship.Y))
	screen.DrawImage(ship.Image, op) //第二个参数可以用于指定坐标相对于原点的偏移：
	// drawimage在图像i上绘制给定的图像。

}
