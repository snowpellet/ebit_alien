package game

import (
	"ebit/bullet"
	config_test "ebit/config"
	"ebit/image"
	img "ebit/image"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	input        *Input                      //输入
	cfg          *config_test.Config         //获取配置信息
	ship         *img.Ship                   //获取图和大小
	bullets      map[*bullet.Bullet]struct{} //管理子弹，子弹十多个
	alien        map[*img.Alien]struct{}     //外星人
	mode         int                         //游戏状态
	failCount    int                         //被外星人碰撞和移除屏幕的外星人
	overMsg      string
	player       *audio.Player  //音频
	audioContent *audio.Context //音频
	counts       int            //记录分数
}

func NewGame() *Game {
	cfg := config_test.LoadConfig()                         //调用加载配置函数
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight) //配置窗口大小
	ebiten.SetWindowTitle(cfg.Title)
	//添加各类效果
	g := &Game{
		input:   &Input{},                                       //键盘输入
		cfg:     cfg,                                            //加载配置信息
		ship:    img.NewShip(cfg.ScreenWidth, cfg.ScreenHeight), //获取图片和大小大小,由于NewShip计算初始坐标需要屏幕尺寸，故增加屏幕宽、高两个参数
		bullets: make(map[*bullet.Bullet]struct{}),              //子弹存储
		alien:   make(map[*img.Alien]struct{}),                  //容纳外星人
		mode:    0,
	}
	// 在游戏一开始就创建一组外星人
	g.CreateAliens()
	g.CreateFonts()
	return g
}

// 在界面打印
func (g *Game) Draw(screen *ebiten.Image) {
	//背景颜色
	screen.Fill(g.cfg.Bgcolor)
	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVADE"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		//绘制飞机的图
		g.ship.Draw(screen, g.cfg)
		// 绘制子弹
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		// 绘制外星人
		for alien := range g.alien {
			alien.Draw(screen)
		}
		msg := fmt.Sprintf("grade=%d", g.counts)
		ebitenutil.DebugPrintAt(screen, msg, 0, 0)

	case ModeOver:
		texts = []string{"", g.overMsg}
	}
	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}

}

// 游戏界面大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

// Game.Update方法调用Input.Update时需要传入飞船对象：
func (g *Game) Update() error {
	//音频搞不来
	// g.audioUpdate()
	switch g.mode {
	case ModeTitle:
		//按下空格变游戏开始
		if IsKeyPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		// 子弹位置更新
		for bullet := range g.bullets {
			// 减去子弹每次移动的距离
			bullet.Y -= bullet.SpeedFactor
		}
		//外星人移动
		for alien := range g.alien { //往下走
			alien.Y += float64(alien.SpendFactor)
		}

		g.input.Update(g)

		g.checkCollision() //检测碰撞
		//子弹超出删除
		for bullet := range g.bullets {
			if bullet.OutOfScreen() {
				delete(g.bullets, bullet)
			}
		}
		for alien1 := range g.alien {
			if alien1.Y >= 480 { //外星人删除
				g.failCount++
				delete(g.alien, alien1)
				continue
			}
			//飞船碰撞外星人
			if bullet.CheckCollisAlien(alien1, g.ship) {
				g.failCount++
				delete(g.alien, alien1)
				continue
			}
		}
		if g.failCount >= 3 {
			g.overMsg = "Game Over!"
		} else if len(g.alien) == 0 {
			g.overMsg = "You Win!"
		}

		if len(g.overMsg) > 0 {
			g.mode = ModeOver
			g.alien = make(map[*image.Alien]struct{})
			g.bullets = make(map[*bullet.Bullet]struct{})
			for alien := range g.alien {
				delete(g.alien, alien)
			}
			for bullet := range g.bullets {
				delete(g.bullets, bullet)
			}
		}
	case ModeOver:
		if IsKeyPressed() {
			g.overMsg = ""
			g.mode = ModeTitle
			g.failCount = 0
			g.ship = img.NewShip(g.cfg.ScreenWidth, g.cfg.ScreenHeight)
			g.CreateAliens()
			g.CreateFonts()
		}
	}

	return nil
}
