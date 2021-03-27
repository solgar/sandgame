package main

import (
	"fmt"
	"log"
	"sandgame/brush"
	"sandgame/particles"
	"sandgame/renderer"
	"sandgame/settings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func init() {
	particles.ClearData()

	for y := 1; y < settings.ScreenHeight-1; y++ {
		for x := 1; x < settings.ScreenWidth-1; x++ {
			*particles.GetDataXY(x, y) = particles.NewSand()
		}
	}

	render = renderer.ReplacePixels
}

type Game struct {
}

var drawStuff = true

var currentBrush = brush.NewDotBrush()

var plusPressed = false
var minusPressed = false

var render renderer.Renderer

var lastFrameTime *time.Time

func (g *Game) Update() error {
	current := time.Now()
	if lastFrameTime != nil {
		fmt.Println("===> whole frame time:", current.Sub(*lastFrameTime))
	}
	lastFrameTime = &current

	mx, my := ebiten.CursorPosition()
	particleX := mx / settings.Scale
	particleY := my / settings.Scale

	_, why := ebiten.Wheel()

	if why > 0 {
		s := currentBrush.GetSize()
		currentBrush.SetSize(s + 1)
	} else if why < 0 {
		s := currentBrush.GetSize()
		currentBrush.SetSize(s - 1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyKPAdd) && !plusPressed {
		s := currentBrush.GetSize()
		currentBrush.SetSize(s + 1)
	}
	plusPressed = ebiten.IsKeyPressed(ebiten.KeyKPAdd)

	if ebiten.IsKeyPressed(ebiten.KeyKPSubtract) && !minusPressed {
		s := currentBrush.GetSize()
		currentBrush.SetSize(s - 1)
	}
	minusPressed = ebiten.IsKeyPressed(ebiten.KeyKPSubtract)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if ebiten.IsKeyPressed(ebiten.KeyControl) {
			currentBrush.PaintType(particleX, particleY, particles.Wood)
		} else {
			currentBrush.PaintType(particleX, particleY, particles.Sand)
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		currentBrush.PaintType(particleX, particleY, particles.Empty)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		currentBrush.PaintType(particleX, particleY, particles.Water)
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		particles.ClearData()
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

	now := time.Now()
	for y := settings.ScreenHeight - 1; y >= 0; y-- {
		for x := 0; x < settings.ScreenWidth; x++ {
			particles.Update(x, y)
		}
	}
	fmt.Println("update time:", time.Now().Sub(now))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	if drawStuff {
		now := time.Now()
		render(screen, particles.GetRawData())
		fmt.Println("render time:", time.Now().Sub(now))
	}

	mx, my := ebiten.CursorPosition()
	particleX := float64(mx / settings.Scale)
	particleY := float64(my / settings.Scale)
	particlePx := int(particleX * settings.Scale)
	particlePy := int(particleY * settings.Scale)

	currentBrush.DrawOutline(screen, particlePx, particlePy)

	msg := fmt.Sprintf("(%d, %d)  FPS: %d", int(particleX), int(particleY), int(ebiten.CurrentFPS()))
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
	ebiten.SetVsyncEnabled(false)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
