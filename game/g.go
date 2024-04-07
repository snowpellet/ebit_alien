package game

import (
	"ebit/bullet"
	img "ebit/image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// 游戏状态
const (
	ModeTitle = iota //游戏界面
	ModeGame         //游戏
	ModeOver         //结束
)

// 字体对象
var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

// 处理字体
func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// 用于添加子弹
func (g *Game) addBullet(bullet *bullet.Bullet) {
	g.bullets[bullet] = struct{}{}
}

// 外星人之间留的空间
func (g *Game) CreateAliens() {
	alien := img.NewAlien(g.cfg)
	//左右各留一个外星人宽度的空间：

	for row := 0; row < 2; row++ { //两行
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		num := r.Intn(alien.Width)
		availableSpaceX := g.cfg.ScreenWidth - 2*num
		//两个外星人之间留一个外星人宽度的空间。所以一行可以创建的外星人的数量为
		numAliens := availableSpaceX / (2 * num)
		// numAliens 几个外星人
		for i := 0; i < numAliens; i++ {
			num1 := r.Intn(alien.Height)
			//两个外星人之间留一个外星人宽度的空间。所以一行可以创建的外星人的数量为
			alien = img.NewAlien(g.cfg)
			//外星人出现之间间隔
			alien.X = float64(num + 2*alien.Width*i)
			alien.Y = float64(float64(alien.Height)*float64(row))*1.5 + float64(num1)/2
			g.AddAlien(alien)
		}
	}
}

// 添加外星人
func (g *Game) AddAlien(alien *img.Alien) {
	g.alien[alien] = struct{}{}
}

// 检测碰撞就删除
func (g *Game) checkCollision() {
	for alien := range g.alien {
		for bullets := range g.bullets {
			if bullet.CheckCollision(bullets, alien) {
				g.counts++
				delete(g.alien, alien)
				delete(g.bullets, bullets)
			}
		}

	}
	if len(g.alien) < 1 {
		g.CreateAliens()
	}
}

// const (
// 	sampleRate     = 48000
// 	bytesPerSample = 4 // 2 channels * 2 bytes (16 bit)

// 	introLengthInSecond = 5
// 	loopLengthInSecond  = 4
// )

// func (g *Game) audioUpdate() {
// 	if g.player != nil {
// 		log.Fatal("player")
// 		return
// 	}
// 	if g.audioContent == nil {
// 		g.audioContent = audio.NewContext(sampleRate)
// 	}
// 	// 打开ogg 音频文件
// 	file, err := ebitenutil.OpenFile("./hr1kc-owol3.ogg")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer file.Close()
// 	// 解码
// 	oggStream, err := ogg.Decode(file)
// 	oggs, err := vorbis.DecodeWithoutResampling(bytes.NewReader(file))
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	s := audio.NewInfiniteLoopWithIntro(oggs, introLengthInSecond*bytesPerSample*sampleRate, loopLengthInSecond*bytesPerSample*sampleRate)
// 	g.player, err = g.audioContent.NewPlayer(s)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	g.player.Play()
// }
