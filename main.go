package main

import (
	"image"
	"log"

	"ebit/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	Game := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("外星人入侵")
	ebiten.SetWindowIcon([]image.Image{})
	if err := ebiten.RunGame(Game); err != nil {
		log.Fatal(err)
	}
}
