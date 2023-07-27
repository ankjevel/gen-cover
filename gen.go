package main

import (
	"bytes"
	"image/color"
	"math"
	"os"
	"strings"

	"github.com/ankjevel/gen-cover/colors"
	"github.com/ankjevel/gen-cover/scene"
	"github.com/lucasb-eyer/go-colorful"
)

var (
	width, height = 400, 400
	offset_y      = int(float64(height) / 2.3)
	offset_x      = int(float64(width) / 6.5)
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

	scene.EachPixel(offset_y, func(x, y int) color.RGBA {
		var (
			xy  = float64(x + y)
			wh  = float64(width + height)
			m   = float64(len(rgbs) - 1)
			i   = (xy / wh) * m
			min = math.Floor(i)
			max = math.Ceil(i)
			p   = i - min
			fi  = int(min)
			ci  = int(math.Min(max, float64(len(rgbs)-1)))
		)

		if fi == ci {
			ci += 1
		}

		if ci == -1 || fi == -1 {
			ci = 0
			fi = 1
		}

		if ci > int(m) {
			ci = int(m)
			fi = ci - 1
		}

		from := &rgbs[fi]
		to := &rgbs[ci]

		if from == nil || to == nil {
			return color.RGBA{0, 0, 0, 255}
		}

		from_rgb := colorful.Color{R: float64(from.R), G: float64(from.G), B: float64(from.B)}
		to_rgb := colorful.Color{R: float64(to.R), G: float64(to.G), B: float64(to.B)}
		blend := from_rgb.BlendHcl(to_rgb, p)

		return color.RGBA{uint8(blend.R), uint8(blend.G), uint8(blend.B), 255}
	})

	scene.AddLabel(offset_x, offset_y, label)

	return scene.Buffer()
}
