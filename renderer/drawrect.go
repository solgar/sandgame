package renderer

import (
	"sandgame/particles"
	"sandgame/settings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawRect(screen *ebiten.Image, data []particles.Particle) {
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
