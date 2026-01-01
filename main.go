package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// TODO: CRIAR BORDAS QUE LIMITAM O PLAYER DE ANDAR PARA FORA DA TELA
// TODO: CRIAR A "MACA" QUE AO COMER O RABO (TAIL) CRESCE

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	snake            *ebiten.Image
	playerX, playerY float64
	tail             int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= 4
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.playerX), float64(g.playerY))
	screen.DrawImage(g.snake, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &Game{}

	g.snake = ebiten.NewImage(10, 10)
	g.snake.Fill(color.White)
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
