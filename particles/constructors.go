package particles

import (
	"image/color"
	"math/rand"
)

func randomWaterColor() color.RGBA {
	max := 256
	min := 230
	rc := uint8(rand.Intn(max-min) + min)
	return color.RGBA{0, 0, rc, 255}
}

func randomWoodColor() color.RGBA {
	return color.RGBA{150, 75, 0, 255}
}

func NewEmpty() Particle {
	return Particle{PType: Empty, Color: color.RGBA{0, 0, 0, 255}}
}

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
	return Particle{PType: Sand,
		Updated: false,
		VelX:    0,
		VelY:    0,
		Color:   color.RGBA{rc, rc, 0, 255}}
}

func NewWater() Particle {
	return Particle{PType: Water,
		Updated: false,
		VelX:    0,
		VelY:    0,
		Color:   randomWaterColor()}
}

func NewWood() Particle {
	return Particle{PType: Wood,
		Updated: false,
		VelX:    0,
		VelY:    0,
		Color:   randomWoodColor()}
}
