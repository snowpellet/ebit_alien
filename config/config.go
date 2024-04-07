package config_test

import (
	"bytes"
	"ebit/resourse"
	"encoding/json"
	"image/color"
	"log"
)

// 配置的结构
type Config struct {
	ScreenWidth       int        `json:"screenWidth"`       //打开窗口大小参数
	ScreenHeight      int        `json:"screenHeight"`      //打开窗口大小参数
	Title             string     `json:"title"`             //标题
	Bgcolor           color.RGBA `json:"bgColor"`           //背景
	ShipSpeedFactor   int        `json:"shipSpeedFactor"`   //每次移动多少像素
	BulletWidth       int        `json:"bulletWidth"`       //子弹宽
	BulletHeight      int        `json:"bulletHeight"`      //子弹高
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"` //子弹移动像素
	BulletBgColor     color.RGBA `json:"bulletBgColor"`     //颜色
	MaxBulletNum      int        `json:"maxbulletnum"`      //最多子弹数量
	BulletInterval    int64      `json:"bulletInterval"`    //子弹间隔
	SpendFactor       int        `json:"spendFactor"`       //外星人移动距离
	TitleFontSize     int
	FontSize          int
	SmallFontSize     int
}

// 加载配置
func LoadConfig() *Config {
	// f, err := os.Open("config/config.json")
	// if err != nil {
	// 	log.Fatalf("os.Open failed: %v\n", err)
	// }

	var cfg Config
	err := json.NewDecoder(bytes.NewReader(resourse.Config)).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
