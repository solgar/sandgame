package renderer

import (
	"image"
	"sandgame/particles"
	"sandgame/settings"

	"github.com/hajimehoshi/ebiten/v2"
)

var screenBuffer *image.RGBA

func ReplacePixels(screen *ebiten.Image, data []particles.Particle) {
	if screenBuffer == nil {
		screenBuffer = image.NewRGBA(image.Rect(0, 0, settings.ResolutionWidth, settings.ResolutionHeight))
	}

	const l = settings.ScreenWidth * settings.ScreenHeight

	for x := 0; x < settings.ScreenWidth; x++ {
		for y := 0; y < settings.ScreenHeight; y++ {
			particleData := particles.GetDataXY(x, y)
			for px := x * settings.Scale; px < (x+1)*settings.Scale; px++ {
				for py := y * settings.Scale; py < (y+1)*settings.Scale; py++ {
					pxIdx := (px + py*settings.ResolutionWidth) * 4
					screenBuffer.Pix[pxIdx] = particleData.Color.R
					screenBuffer.Pix[pxIdx+1] = particleData.Color.G
					screenBuffer.Pix[pxIdx+2] = particleData.Color.B
					screenBuffer.Pix[pxIdx+3] = particleData.Color.A
				}
			}

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

	screen.ReplacePixels(screenBuffer.Pix)
}
