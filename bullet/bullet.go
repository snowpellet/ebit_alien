package bullet

// 绘制子弹()
import (
	config_test "ebit/config"
	pic "ebit/image"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Image       *ebiten.Image
	Width       int
	Height      int
	X           float64
	Y           float64
	SpeedFactor float64 //移动速度
}

func NewBullet(cfg *config_test.Config, ship *pic.Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	/*
		Rect是矩形(Pt(x0, y0)， Pt(x1, y1)}的简写。
		返回的矩形在必要时交换了最小和最大坐标，以使其格式良好。

	*/
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletBgColor)
	/*
		NewlmageWithOptions返回一个带有给定边界和选项的空图像。
		如果宽度或高度小于1或大于与设备相关的最大大小，NewlmageWithOptions会出现恐慌。
		渲染的起始位置是给定边界的(0,0)。
		如果在NewlmageOptions创建的新图像上调用Drawlmage，例如，缩放和旋转的中心是(0,0)，那可能不是左上角的位置。

	*/
	return &Bullet{
		Image:       img,
		Width:       cfg.BulletWidth,
		Height:      cfg.BulletHeight,
		X:           float64(cfg.ScreenWidth) - (float64(ship.X) + float64(ship.Width)/2), //y轴子弹是从飞船头部发送，横坐标等于飞船中心的横坐标，(改左上角的纵坐标=屏幕高度-飞船高-子弹高。)
		Y:           float64(cfg.ScreenHeight - ship.Y - cfg.BulletHeight),                //
		SpeedFactor: cfg.BulletSpeedFactor,
	}
}

// 绘制子弹，显示处理
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bullet.X), float64(bullet.Y)) //Translate(tx,ty)平移矩阵
	// GeoM是要绘制的几何矩阵。
	//默认(0)值是identity，它在(0,0)处绘制图像。
	screen.DrawImage(bullet.Image, op)
	// 在图像 i 上绘制给定的图像。
}

// 判断是否处于屏幕外的方法：
func (bullet *Bullet) OutOfScreen() bool {
	return bullet.Y < float64(-bullet.Height)

}
