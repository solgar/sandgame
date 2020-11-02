package main

import (
	"fmt"
	"log"
	"sandgame/particles"
	"sandgame/settings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	for x := 0; x < settings.ScreenWidth; x++ {
		*particles.GetDataXY(x, 0) = particles.NewRock()
		*particles.GetDataXY(x, settings.ScreenHeight-1) = particles.NewRock()
	}

	for y := 0; y < settings.ScreenHeight; y++ {
		*particles.GetDataXY(0, y) = particles.NewRock()
		*particles.GetDataXY(settings.ScreenWidth-1, y) = particles.NewRock()
	}
}

type Game struct {
}

var drawStuff = true

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	particleX := mx / settings.Scale
	particleY := my / settings.Scale

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			(*particles.GetDataXY(particleX, particleY)) = particles.NewWood()
		} else {
			(*particles.GetDataXY(particleX, particleY)) = particles.NewSand()
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		(*particles.GetDataXY(particleX-1, particleY-1)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX, particleY-1)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX+1, particleY-1)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX-1, particleY)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX, particleY)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX+1, particleY)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX-1, particleY+1)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX, particleY+1)) = particles.Particle{PType: particles.Empty}
		(*particles.GetDataXY(particleX+1, particleY+1)) = particles.Particle{PType: particles.Empty}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		(*particles.GetDataXY(particleX, particleY)) = particles.NewWater()
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		drawStuff = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		drawStuff = false
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

	if drawStuff {
		for x := 0; x < settings.ScreenWidth; x++ {
			for y := settings.ScreenHeight - 1; y >= 0; y-- {
				particleData := particles.GetDataXY(x, y)
				if particleData.PType != particles.Empty {
					px := float64(x * settings.Scale)
					py := float64(y * settings.Scale)
					ebitenutil.DrawRect(screen, px, py, settings.Scale, settings.Scale, particleData.Color)
				}
			}
		}
	}

	mx, my := ebiten.CursorPosition()
	msg := fmt.Sprintf("(%d, %d)", mx, my)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.ScreenWidth * settings.Scale, settings.ScreenHeight * settings.Scale
}

func main() {
	// ebiten.SetMaxTPS(15)
	ebiten.SetWindowSize(settings.ScreenWidth*settings.Scale, settings.ScreenHeight*settings.Scale)
	ebiten.SetWindowTitle("sandgame")
	ebiten.SetFullscreen(true)
	ebiten.SetVsyncEnabled(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
