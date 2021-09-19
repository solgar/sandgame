package particles

import (
	"image/color"
)

type ParticleType uint8

const (
	Empty = iota
	Smoke
	WaterVapor
	Fire
	Water
	Acid
	Oil
	Wood
	Sand
	Rock
	OOB
)

type Particle struct {
	PType   ParticleType
	VelX    float64
	VelY    float64
	Updated bool
	Color   color.RGBA
}

func (p *Particle) IsGas() bool {
	return p.PType == Smoke || p.PType == WaterVapor
}

func (p *Particle) IsLiquid() bool {
	return p.PType == Water || p.PType == Acid || p.PType == Oil
}

func (p *Particle) IsSolid() bool {
	return p.PType >= Wood && p.PType < OOB
}

func (p *Particle) IsLoose() bool {
	// no loose particles atm
	return false
}
