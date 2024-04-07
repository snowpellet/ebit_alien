package bullet

import (
	config_test "ebit/config"
	"ebit/image"
)

// 检测子弹碰撞外星人
func CheckCollision(bullet *Bullet, alien *image.Alien) bool {
	alienTop, alienLeft := alien.Y, alien.X //得到外星人图的头位置和右边，头图的高度，就能的到图的尾，也就是子弹碰撞的位置
	alienBottom, alienRight := alien.Y+float64(alien.Height), alien.X+float64(alien.Width)
	// 左上角

	x, y := bullet.X, bullet.Y //外星人移动y+,子弹往上是y-
	if y < alienBottom && y > alienTop && x > alienLeft && x < alienRight {
		return true
	}
	x, y = bullet.X+float64(bullet.Width), bullet.Y
	if y < alienBottom && y > alienTop && x > alienLeft && x < alienRight {
		return true
	}
	// x, y = bullet.X, bullet.Y
	// if y <  alienBottom) && y >  alienTop) && x >  alienLeft) && x <  alienRight) {
	// 	return true
	// }

	return false
}

// 检测与飞船发生碰撞
func CheckCollisAlien(alien *image.Alien, ship *image.Ship) bool {
	cfg := config_test.LoadConfig()
	top, left := alien.Y, alien.X //得到外星人图的头位置和右边，头图的高度，就能的到图的尾，也就是子弹碰撞的位置
	bottom, right := alien.Y+float64(alien.Height), alien.X+float64(alien.Width)
	x, y := ship.X, ship.Y
	// fmt.Println(x, y, cfg.ScreenHeight-int(top), cfg.ScreenHeight-int(bottom), cfg.ScreenWidth-int(left), cfg.ScreenWidth-int(right))
	// 左上角
	if (y < cfg.ScreenHeight-int(top) || y > cfg.ScreenHeight-int(top)) && y > cfg.ScreenHeight-int(bottom) && x < cfg.ScreenWidth-int(left) && x > cfg.ScreenWidth-int(right) {
		// fmt.Println(1)
		return true
	}

	// 右上角
	x, y = ship.X+ship.Width, ship.Y
	if (y < cfg.ScreenHeight-int(top) || y > cfg.ScreenHeight-int(top)) && y > cfg.ScreenHeight-int(bottom) && x < cfg.ScreenWidth-int(left) && x > cfg.ScreenWidth-int(right) {
		// fmt.Println(2)
		return true
	}

	// 左下角
	x, y = ship.X, ship.Y+ship.Height
	if y < cfg.ScreenHeight-int(top) && y >= cfg.ScreenHeight-int(bottom) && x < cfg.ScreenWidth-int(left) && x > cfg.ScreenWidth-int(right) {
		// fmt.Println(3)
		return true
	}

	// 右下角
	x, y = ship.X+ship.Width, ship.Y+ship.Width
	if y < cfg.ScreenHeight-int(top) && y >= cfg.ScreenHeight-int(bottom) && x < cfg.ScreenWidth-int(left) && x > cfg.ScreenWidth-int(right) {
		// fmt.Println(4)
		return true
	}
	return false
}
