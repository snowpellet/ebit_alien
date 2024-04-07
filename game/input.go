package game

import (
	"ebit/bullet"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	lastBulletTime time.Time //记录上次发射子弹的时间
}

// 然后我们在Input的Update方法中根据按下的是左方向键还是右方向键来更新飞船的坐标：
func (i *Input) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) { //空格发射子弹
		// g.cfg.MaxBulletNum()子弹数量
		if g.cfg.MaxBulletNum > len(g.bullets) && time.Now().Sub(i.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
			bullet := bullet.NewBullet(g.cfg, g.ship)
			g.addBullet(bullet)
			i.lastBulletTime = time.Now()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.ship.X -= g.cfg.ShipSpeedFactor
		if g.ship.X < -g.ship.Width/2 {
			g.ship.X = -g.ship.Width / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.ship.X += g.cfg.ShipSpeedFactor
		if g.ship.X > g.cfg.ScreenWidth-g.ship.Width/2 {
			g.ship.X = g.cfg.ScreenWidth - g.ship.Width/2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) { //下
		// fmt.Println(g.ship.Height)
		g.ship.Y -= g.cfg.ShipSpeedFactor
		if g.ship.Y <= 0 {
			g.ship.Y = 48
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) { //上+
		// fmt.Println(g.ship.Height, g.cfg.ScreenHeight)
		g.ship.Y += g.cfg.ShipSpeedFactor
		if g.ship.Y >= g.cfg.ScreenHeight {
			g.ship.Y = 48
		}
	}
}

// 判断是否按下空格和左键
func IsKeyPressed() bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		return true
	}
	//左键
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}
