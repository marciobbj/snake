package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

// TODO: CRIAR A "MACA" QUE AO COMER O RABO (TAIL) CRESCE

const (
	screenWidth     = 640
	screenHeight    = 480
	limitHorizontal = screenWidth - 20
	limitVertical   = screenHeight - 20
	debug           = true
)

type Game struct {
	snake            *ebiten.Image
	playerX, playerY float64
	tail             int
	limits           *ebiten.Image
	rx, ry           float32
	sh, sw           int
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
	if g.playerX < 10.0 || g.playerX > limitHorizontal {
		//TODO mata tudo e reseta
		g.playerX = float64(screenWidth / 2)
		g.playerY = float64(screenHeight / 2)
		// TODO criar algo como resetaftercolition()
	}
	if g.playerY < 10 || g.playerY > limitVertical {
		g.playerX = float64(screenWidth / 2)
		g.playerY = float64(screenHeight / 2)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Fundo preto

	g.sw = screen.Bounds().Dx()
	g.sh = screen.Bounds().Dy()
	g.rx = float32(g.sw)/2 - float32(limitHorizontal)/2
	g.ry = float32(g.sh)/2 - float32(limitVertical)/2
	vector.StrokeRect(screen, g.rx, g.ry, float32(limitHorizontal), float32(limitVertical), 2, color.White, true)

	snakeOpts := &ebiten.DrawImageOptions{}
	snakeOpts.GeoM.Translate(float64(g.playerX), float64(g.playerY))
	screen.DrawImage(g.snake, snakeOpts)

	if debug {
		ebitenutil.DebugPrintAt(screen, "Player Position: "+fmt.Sprintf("%0.2f", g.playerX)+" X "+fmt.Sprintf("%0.2f", g.playerY), limitHorizontal-200, limitVertical-30)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func NewGame() *Game {
	g := &Game{}
	//centraliza o spawn do ponto no centro da tela
	g.playerX = float64(screenWidth / 2)
	g.playerY = float64(screenHeight / 2)

	g.snake = ebiten.NewImage(10, 10)
	g.snake.Fill(color.RGBA{255, 0, 0, 255}) // Ponto vermelho para destacar do fundo

	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame()
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
