package brush

import (
	"sandgame/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

type Brush interface {
	DrawOutline(dst *ebiten.Image, px, py int)
	PaintParticle(x, y int, p *particles.Particle)
	PaintType(x, y int, pType particles.ParticleType)

	GetSize() int
	SetSize(int)
}

type brush struct {
	drawOutlineF   func(*ebiten.Image, *brush, int, int)
	paintParticleF func(*brush, int, int, *particles.Particle)
	paintPType     func(*brush, int, int, particles.ParticleType)

	size int
}

func (b *brush) DrawOutline(dst *ebiten.Image, px, py int) {
	b.drawOutlineF(dst, b, px, py)
}

func (b *brush) PaintParticle(x, y int, p *particles.Particle) {
	b.paintParticleF(b, x, y, p)
}

func (b *brush) PaintType(x, y int, pType particles.ParticleType) {
	b.paintPType(b, x, y, pType)
}

func (b *brush) GetSize() int {
	return b.size
}

func (b *brush) SetSize(newSize int) {
	if newSize < 1 {
		return
	}
	b.size = newSize
}
