package brush

import (
	"image/color"
	"sandgame/particles"
	"sandgame/settings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewDotBrush() Brush {
	return &brush{drawOutlineF: drawDotOutline, paintParticleF: dotPaintParticle, paintPType: dotPaintType, size: 2}
}

func drawDotOutline(dst *ebiten.Image, b *brush, x, y int) {
	px := float64(x)
	py := float64(y)
	halfParticle := float64(settings.Scale) / 2

	brushWidth := float64(settings.Scale * b.size)
	brushHeight := float64(settings.Scale * b.size)

	ebitenutil.DrawRect(dst, px-halfParticle,
		py-halfParticle,
		brushWidth+2*halfParticle,
		halfParticle,
		color.White)
	ebitenutil.DrawRect(dst, px-halfParticle,
		py+brushHeight,
		brushWidth+2*halfParticle,
		halfParticle,
		color.White)
	ebitenutil.DrawRect(dst, px-halfParticle,
		py-halfParticle,
		halfParticle,
		brushHeight+2*halfParticle,
		color.White)
	ebitenutil.DrawRect(dst, px+brushWidth,
		py-halfParticle,
		halfParticle,
		brushHeight+2*halfParticle,
		color.White)
}

func dotPaintParticle(b *brush, x, y int, p *particles.Particle) {
	dotPaintType(b, x, y, p.PType)
}

func dotPaintType(b *brush, x, y int, pType particles.ParticleType) {
	for ix := x; ix < x+b.size; ix++ {
		for iy := y; iy < y+b.size; iy++ {
			if particles.IsInBoundsDataXY(ix, iy) {
				switch pType {
				case particles.Sand:
					*particles.GetDataXY(ix, iy) = particles.NewSand()
				case particles.Water:
					*particles.GetDataXY(ix, iy) = particles.NewWater()
				case particles.Wood:
					*particles.GetDataXY(ix, iy) = particles.NewWood()
				case particles.Empty:
					*particles.GetDataXY(ix, iy) = particles.NewEmpty()
				}
			}
		}
	}
}
