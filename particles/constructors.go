package particles

import (
	"image/color"
	"math/rand"
)

func NewRock() Particle {
	max := 150
	min := 120
	rc := uint8(rand.Intn(max-min) + min)
	return Particle{PType: Rock,
		Updated: false,
		VelX:    0,
		VelY:    0,
		Color:   color.RGBA{rc, rc, rc, 255}}
}

func NewSand() Particle {
	max := 220
	min := 180
	rc := uint8(rand.Intn(max-min) + min)
	return Particle{PType: Water,
		Updated: false,
		VelX:    0,
		VelY:    0,
		Color:   color.RGBA{rc, rc, 0, 255}}
}

func NewWater() Particle {
	max := 256
	min := 230
	rc := uint8(rand.Intn(max-min) + min)
	return Particle{PType: Water,
		Updated: false,
		VelX:    0,
		VelY:    0,
		Color:   color.RGBA{0, 0, rc, 255}}
}
