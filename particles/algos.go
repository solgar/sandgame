package particles

import (
	"image/color"
	"math/rand"
	"sandgame/settings"
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

func ClearData() {
	for x := 0; x < settings.ScreenWidth; x++ {
		for y := 0; y < settings.ScreenHeight; y++ {
			*GetDataXY(x, y) = NewEmpty()
		}
	}

	for x := 0; x < settings.ScreenWidth; x++ {
		*GetDataXY(x, 0) = NewWood()
		*GetDataXY(x, settings.ScreenHeight-1) = NewWood()
	}

	for y := 0; y < settings.ScreenHeight; y++ {
		*GetDataXY(0, y) = NewWood()
		*GetDataXY(settings.ScreenWidth-1, y) = NewWood()
	}
}

func (p *Particle) IsGas() bool {
	return p.PType >= Smoke && p.PType <= WaterVapor
}

func (p *Particle) IsLiquid() bool {
	return p.PType >= Water && p.PType <= Oil
}

func (p *Particle) IsSolid() bool {
	return p.PType >= Wood && p.PType <= Rock
}

var oobParticle = &Particle{PType: OOB}

var data []Particle = make([]Particle, settings.ScreenWidth*settings.ScreenHeight)

var randSeed int64 = 0

func init() {
	rand.Seed(randSeed)
}

func IsInBoundsDataXY(x, y int) bool {
	return x >= 0 && x < settings.ScreenWidth && y >= 0 && y < settings.ScreenHeight
}

func GetRawData() []Particle {
	return data
}

func GetDataXY(x, y int) *Particle {
	if IsInBoundsDataXY(x, y) {
		return &data[x+y*settings.ScreenWidth]
	}
	return oobParticle
}

func SetDataXY(x, y int, p *Particle) {
	data[x+y*settings.ScreenWidth] = *p
}

func SwapDataXY(x1, y1, x2, y2 int) {
	p1 := GetDataXY(x1, y1)
	p2 := GetDataXY(x2, y2)
	SwapParticles(p1, p2)
}

func SwapParticles(p1, p2 *Particle) {
	p1c := *p1
	*p1 = *p2
	*p2 = p1c
}

func Update(x, y int) {
	p := GetDataXY(x, y)
	if !p.Updated {
		p.Updated = true
		switch p.PType {
		case Empty:
			updateEmpty(x, y, p)
		case Sand:
			updateSand(x, y, p)
		case Water:
			p.Color = randomWaterColor()
			updateWater(x, y, p)
		}
	}
}

func updateEmpty(x, y int, p *Particle) {
	if GetDataXY(x-2, y).PType == Water &&
		GetDataXY(x-1, y).PType == Water &&
		GetDataXY(x+1, y).PType == Water &&
		GetDataXY(x+2, y).PType == Water &&
		GetDataXY(x-2, y+1).PType == Water &&
		GetDataXY(x-1, y+1).PType == Water &&
		GetDataXY(x, y+1).PType == Water &&
		GetDataXY(x+1, y+1).PType == Water &&
		GetDataXY(x+2, y+1).PType == Water {
		*p = NewWater()
	}
}

func updateSand(x, y int, p *Particle) {
	bottomP := GetDataXY(x, y+1)
	if bottomP.PType == Empty {
		SwapDataXY(x, y, x, y+1)
	} else if bottomP.IsSolid() {
		toLD := GetDataXY(x-1, y+1)
		toRD := GetDataXY(x+1, y+1)
		if toLD.PType == Empty && toRD.PType == Empty {
			direction := rand.Int() % 2
			if direction == 0 {
				SwapDataXY(x, y, x-1, y+1)
			} else {
				SwapDataXY(x, y, x+1, y+1)
			}
		} else if toLD.PType == Empty {
			SwapDataXY(x, y, x-1, y+1)
		} else if toRD.PType == Empty {
			SwapDataXY(x, y, x+1, y+1)
		} else if toLD.IsLiquid() && toRD.IsLiquid() {
			toL := GetDataXY(x-1, y)
			toR := GetDataXY(x+1, y)
			if toL.IsLiquid() && toR.IsLiquid() {
				direction := rand.Int() % 2
				if direction == 0 {
					SwapParticles(p, toLD)
				} else {
					SwapParticles(p, toRD)
				}
			}
		}
	} else if bottomP.IsLiquid() {
		toL := GetDataXY(x-1, y)
		toR := GetDataXY(x+1, y)
		direction := rand.Int() % 2
		if toL.IsLiquid() || toR.IsLiquid() {
			toLD := GetDataXY(x-1, y+1)
			toRD := GetDataXY(x+1, y+1)
			direction := rand.Int() % 3
			if direction == 0 && toLD.IsLiquid() {
				SwapParticles(p, toLD)
			} else if direction == 1 && toRD.IsLiquid() {
				SwapParticles(p, toRD)
			} else {
				SwapParticles(p, bottomP)
			}
		} else if direction == 0 && toL.PType == Empty {
			*toL = *bottomP
			*bottomP = *p
			*p = NewEmpty()
		} else if toR.PType == Empty {
			*toR = *bottomP
			*bottomP = *p
			*p = NewEmpty()
		} else {
			// *bottomP = *p
			// *p = NewEmpty()
			SwapParticles(p, bottomP)
		}
	}
}

func updateWater(x, y int, p *Particle) {
	if GetDataXY(x, y+1).PType == Empty {
		SwapDataXY(x, y, x, y+1)
	} else {
		isBottomLeftEmpty := GetDataXY(x-1, y+1).PType == Empty
		isBottomRightEmpty := GetDataXY(x+1, y+1).PType == Empty
		if isBottomLeftEmpty && isBottomRightEmpty {
			if (rand.Int() % 2) == 0 {
				SwapDataXY(x, y, x-1, y+1)
			} else {
				SwapDataXY(x, y, x+1, y+1)
			}
		} else if isBottomLeftEmpty {
			SwapDataXY(x, y, x-1, y+1)
		} else if isBottomRightEmpty {
			SwapDataXY(x, y, x+1, y+1)
		} else {
			if p.VelX == 0 {
				if (rand.Int() % 2) == 0 {
					p.VelX = -1
				} else {
					p.VelX = 1
				}
			}
			neighbour := GetDataXY(x+int(p.VelX), y)
			if neighbour.PType == Empty {
				SwapParticles(p, neighbour)
			} else {
				p.VelX *= -1
			}
		}
	}
}
