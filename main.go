package main

import (
	"fmt"
	"image/color"
	"log"
	"sandgame/particles"
	"sandgame/settings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	for x := 0; x < settings.ScreenWidth; x++ {
		particles.GetDataXY(x, 0).PType = particles.Solid
		particles.GetDataXY(x, settings.ScreenHeight-1).PType = particles.Solid
	}

	for y := 0; y < settings.ScreenHeight; y++ {
		particles.GetDataXY(0, y).PType = particles.Solid
		particles.GetDataXY(settings.ScreenWidth-1, y).PType = particles.Solid
	}
}

type Game struct {
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		particles.GetDataXY(mx, my).PType = particles.Sand
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		particles.GetDataXY(mx, my).PType = particles.Water
	}

	for y := settings.ScreenHeight - 1; y >= 0; y-- {
		for x := 0; x < settings.ScreenWidth; x++ {
			particles.GetDataXY(x, y).Updated = false
		}
	}

	for y := settings.ScreenHeight - 1; y >= 0; y-- {
		for x := 0; x < settings.ScreenWidth; x++ {
			particles.Update(x, y)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	for x := 0; x < settings.ScreenWidth; x++ {
		for y := settings.ScreenHeight - 1; y >= 0; y-- {
			particleData := particles.GetDataXY(x, y)
			switch particleData.PType {
			case particles.Sand:
				ebitenutil.DrawRect(screen, float64(x), float64(y), 1, 1, color.RGBA{R: 255, G: 255, B: 0, A: 255})
			case particles.Water:
				ebitenutil.DrawRect(screen, float64(x), float64(y), 1, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255})
			}
		}
	}

	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("(%d, %d)", mx, my)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.ScreenWidth, settings.ScreenHeight
}

func main() {
	// ebiten.SetMaxTPS(15)
	ebiten.SetWindowSize(settings.ScreenWidth*settings.Scale, settings.ScreenHeight*settings.Scale)
	ebiten.SetWindowTitle("Paint (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
