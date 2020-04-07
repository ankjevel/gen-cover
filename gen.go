package main

import (
	"bytes"
	"image/color"
	"math"
	"os"
	"strings"

	"github.com/Iteam1337/gen-cover/colors"
	"github.com/Iteam1337/gen-cover/scene"
	"github.com/lucasb-eyer/go-colorful"
)

var (
	width, height = 400, 400
	offsetY       = int(float64(height) / 2.3)
	offsetX       = int(float64(width) / 6.5)
)

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value != "" {
		return value
	}

	return fallback
}

func gen(label string) (bytes.Buffer, error) {
	scene := scene.New(width, height)
	rgbs := colors.Get(label)

	scene.EachPixel(offsetY, func(x, y int) color.RGBA {
		var (
			xy = float64(x + y)
			wh = float64(width + height)
			i  = (xy / wh) * float64(len(rgbs))
			f  = math.Floor(i)
			c  = math.Ceil(i)
			p  = i - f
		)

		pc := &rgbs[int(f)]
		nc := &rgbs[int(math.Min(c, float64(len(rgbs)-1)))]

		if pc == nil || nc == nil {
			return color.RGBA{0, 0, 0, 255}
		}

		pcrgb := colorful.Color{R: float64(pc.R), G: float64(pc.G), B: float64(pc.B)}
		ncrgb := colorful.Color{R: float64(nc.R), G: float64(nc.G), B: float64(nc.B)}

		blend := pcrgb.BlendHcl(ncrgb, p)

		return color.RGBA{uint8(blend.R), uint8(blend.G), uint8(blend.B), 255}
	})

	scene.AddLabel(offsetX, offsetY, label)

	return scene.Buffer()
}
