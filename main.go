package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 380
	gridSize        = 20
	margin          = 60
	limitHorizontal = screenWidth - (margin * 2)
	limitVertical   = screenHeight - (margin * 2)
)

type SnakePart struct {
	X, Y float64
}

type Game struct {
	snake      []SnakePart
	apple      SnakePart
	dirX, dirY float64
	timer      int
	moveSpeed  int
	score      int // Novo campo para o placar
}

func NewGame() *Game {
	g := &Game{
		snake: []SnakePart{
			{X: margin + (gridSize * 5), Y: margin + (gridSize * 5)},
		},
		dirX:      gridSize,
		dirY:      0,
		moveSpeed: 8,
		score:     0,
	}
	g.spawnApple()
	return g
}

func (g *Game) spawnApple() {
	cols := limitHorizontal / gridSize
	rows := limitVertical / gridSize
	g.apple.X = float64(rand.IntN(cols))*gridSize + margin
	g.apple.Y = float64(rand.IntN(rows))*gridSize + margin
}

func (g *Game) Update() error {
	// Controles
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && g.dirY == 0 {
		g.dirX, g.dirY = 0, -gridSize
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && g.dirY == 0 {
		g.dirX, g.dirY = 0, gridSize
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && g.dirX == 0 {
		g.dirX, g.dirY = -gridSize, 0
	} else if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && g.dirX == 0 {
		g.dirX, g.dirY = gridSize, 0
	}

	g.timer++
	if g.timer >= g.moveSpeed {
		g.timer = 0

		oldHead := g.snake[0]
		newHead := SnakePart{X: oldHead.X + g.dirX, Y: oldHead.Y + g.dirY}

		// Lógica de comer e crescer
		if newHead.X == g.apple.X && newHead.Y == g.apple.Y {
			g.snake = append([]SnakePart{newHead}, g.snake...)
			g.score += 10 // Aumenta o placar
			g.spawnApple()
		} else {
			g.snake = append([]SnakePart{newHead}, g.snake[:len(g.snake)-1]...)
		}

		// Morte por colisão com bordas
		head := g.snake[0]
		if head.X < margin || head.X >= margin+limitHorizontal || head.Y < margin || head.Y >= margin+limitVertical {
			*g = *NewGame()
		}

		// Morte por colisão com o próprio corpo (Auto-canibalismo)
		for i := 1; i < len(g.snake); i++ {
			if head.X == g.snake[i].X && head.Y == g.snake[i].Y {
				*g = *NewGame()
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// --- INTERFACE E LOGS ---

	// 1. Placar (Score) no topo
	scoreMsg := fmt.Sprintf("PONTUAÇÃO: %d", g.score)
	ebitenutil.DebugPrintAt(screen, scoreMsg, screenWidth/2-40, 20)

	// 2. Logs de Debug (Melhorados)
	statusMsg := fmt.Sprintf(
		"Cobra: %d partes\nCabeça: (%.0f, %.0f)\nMaçã: (%.0f, %.0f)\nFPS: %.2f",
		len(g.snake),
		g.snake[0].X, g.snake[0].Y,
		g.apple.X, g.apple.Y,
		ebiten.ActualFPS(),
	)
	ebitenutil.DebugPrintAt(screen, statusMsg, 15, 15)

	// --- ELEMENTOS DO JOGO ---

	// Moldura
	vector.StrokeRect(screen, float32(margin), float32(margin), float32(limitHorizontal), float32(limitVertical), 2, color.White, true)

	// Maçã
	vector.DrawFilledRect(screen, float32(g.apple.X), float32(g.apple.Y), gridSize-2, gridSize-2, color.RGBA{255, 0, 0, 255}, true)

	// Cobra
	for i, part := range g.snake {
		c := color.RGBA{255, 255, 255, 255}
		if i == 0 {
			c = color.RGBA{150, 255, 150, 255} // Cabeça verde clara
		}
		vector.DrawFilledRect(screen, float32(part.X), float32(part.Y), gridSize-2, gridSize-2, c, true)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Go - Refatorado")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}

