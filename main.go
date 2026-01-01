package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand/v2"
)

// TODO: CRIAR A "MACA" QUE AO COMER O RABO (TAIL) CRESCE

const (
	screenWidth     = 640
	screenHeight    = 480
	limitHorizontal = screenWidth - 20
	limitVertical   = screenHeight - 20
	debug           = true
)

type ApplePos struct {
	x int
	y int
}

type Game struct {
	snake            *ebiten.Image
	playerX, playerY float64
	tail             int
	limits           *ebiten.Image
	rx, ry           float32
	sh, sw           int
	apple_spot       *ApplePos
	apple            *ebiten.Image
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
	g.SpawnApple(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *Game) SpawnApple(screen *ebiten.Image) {
	appleOpts := &ebiten.DrawImageOptions{}
	appleOpts.GeoM.Translate(float64(g.apple_spot.x), float64(g.apple_spot.y))
	screen.DrawImage(g.apple, appleOpts)
}

func GenerateRandomApplePosition() ApplePos {
	x_min_pos := 15
	x_max_pos := limitHorizontal - 15
	y_min_pos := 10
	y_max_pos := limitVertical - 15

	rangeSizeX := x_max_pos - x_min_pos + 1
	rangeSizeY := y_max_pos - y_min_pos + 1

	randCoordX := rand.IntN(rangeSizeX) + x_min_pos
	randCoordY := rand.IntN(rangeSizeY) + y_min_pos

	randOutput := ApplePos{x: randCoordX, y: randCoordY}
	return randOutput
}

func NewGame() *Game {
	g := &Game{}
	//centraliza o spawn do ponto no centro da tela
	g.playerX = float64(screenWidth / 2)
	g.playerY = float64(screenHeight / 2)

	g.snake = ebiten.NewImage(10, 10)
	g.snake.Fill(color.RGBA{179, 199, 57, 255})

	g.apple = ebiten.NewImage(5, 5)
	g.apple.Fill(color.RGBA{255, 0, 0, 255})
	apple_pos := GenerateRandomApplePosition()
	g.apple_spot = &apple_pos
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
