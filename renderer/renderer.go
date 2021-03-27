package renderer

import (
	"sandgame/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer func(screen *ebiten.Image, data []particles.Particle)
