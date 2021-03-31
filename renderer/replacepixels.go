package renderer

import (
	"image"
	"sandgame/particles"
	"sandgame/settings"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenBuffer    *image.RGBA
	screenBufferImg *ebiten.Image
)

func ReplacePixels(screen *ebiten.Image, data []particles.Particle) {
	if screenBuffer == nil {
		screenBuffer = image.NewRGBA(image.Rect(0, 0, settings.ScreenWidth, settings.ScreenHeight))
	}
	if screenBufferImg == nil {
		screenBufferImg = ebiten.NewImage(settings.ScreenWidth, settings.ScreenHeight)
	}

	const l = settings.ScreenWidth * settings.ScreenHeight

	for x := 0; x < settings.ScreenWidth; x++ {
		for y := 0; y < settings.ScreenHeight; y++ {
			particleData := particles.GetDataXY(x, y)
			pxIdx := (x + y*settings.ScreenWidth) * 4
			screenBuffer.Pix[pxIdx] = particleData.Color.R
			screenBuffer.Pix[pxIdx+1] = particleData.Color.G
			screenBuffer.Pix[pxIdx+2] = particleData.Color.B
			screenBuffer.Pix[pxIdx+3] = particleData.Color.A
		}
	}

	/*
		for i := 0; i < l; i++ {
			particleData := data[i]
			if particleData.PType != particles.Empty {
				pxIdx := 4 * i
				screenBuffer.Pix[pxIdx] = particleData.Color.R
				screenBuffer.Pix[pxIdx+1] = particleData.Color.G
				screenBuffer.Pix[pxIdx+2] = particleData.Color.B
				screenBuffer.Pix[pxIdx+3] = particleData.Color.A
				// for r := 0; r < settings.Scale; r++ {
				// 	for c := 0; c < settings.Scale; c++ {
				// 		pxIdx := 4*i*settings.Scale + c + r*settings.ScreenWidth
				// 		fmt.Println("r:", i, "c:", c, "pxIdx:", pxIdx)
				// 		screenBuffer.Pix[pxIdx] = particleData.Color.R
				// 		screenBuffer.Pix[pxIdx+1] = particleData.Color.G
				// 		screenBuffer.Pix[pxIdx+2] = particleData.Color.B
				// 		screenBuffer.Pix[pxIdx+3] = particleData.Color.A
				// 	}
				// }
				// ebitenutil.DrawRect(screen, px, py, settings.Scale, settings.Scale, particleData.Color)

			}
		}
	*/

	screenBufferImg.ReplacePixels(screenBuffer.Pix)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(settings.Factor, settings.Factor)
	screen.DrawImage(screenBufferImg, op)
}
