package particles

import (
	"image/color"
	"math/rand"
	"sandgame/settings"
)

type ParticleType uint8

const (
	Empty = iota
	Sand
	Water
	Fire
	WaterVapor
	Smoke
	Rock
	Solid
	OOB
)

type Particle struct {
	PType   ParticleType
	VelX    float64
	VelY    float64
	Updated bool
	Color   color.RGBA
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
	} else if bottomP.PType != OOB {
		isBottomLeftEmpty := GetDataXY(x-1, y+1).PType == Empty
		isBottomRightEmpty := GetDataXY(x+1, y+1).PType == Empty
		if isBottomLeftEmpty && isBottomRightEmpty {
			direction := rand.Int() % 2
			if direction == 0 {
				SwapDataXY(x, y, x-1, y+1)
			} else {
				SwapDataXY(x, y, x+1, y+1)
			}
		} else if isBottomLeftEmpty {
			SwapDataXY(x, y, x-1, y+1)
		} else if isBottomRightEmpty {
			SwapDataXY(x, y, x+1, y+1)
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
