package scene

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"

	"github.com/ankjevel/gen-cover/fonts"
	"github.com/ankjevel/gen-cover/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	noise image.Image
)

type Scene struct {
	Width, Height int
	Img           *image.RGBA
}

func (s *Scene) AddLabel(x, y int, label string) {
	font_size := float64(fonts.FaceMainSize)
	dot_x := fixed.Int26_6(font_size * 1.4 * 64)
	size := s.Img.Bounds().Size()
	drawer := func(y int, label string, face font.Face) {
		drawer := &font.Drawer{
			Dst:  s.Img,
			Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
			Face: face,
			Dot: fixed.Point26_6{
				X: dot_x,
				Y: fixed.Int26_6(y),
			},
		}
		drawer.DrawString(label)
	}

	sy := float64(size.Y)
	if utils.Config.Title != "" {
		drawer(int(sy-(font_size*6.6))*64, utils.Config.Title, fonts.FaceSub)
	}

	drawer(int(sy-(font_size*4.8))*64, label, fonts.FaceMain)
}

func (s *Scene) EachPixel(offset int, colorFunction func(int, int) color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.Img.Set(x, y, colorFunction(x, y))
		}
	}

	bounds := s.Img.Bounds()

	m := image.NewRGBA(bounds)
	draw.Draw(m, bounds, s.Img, image.Point{0, 0}, draw.Src)
	draw.Draw(m, bounds, noise, image.Point{0, 0}, draw.Over)

	s.Img = m
}

func (s *Scene) Buffer() (buf bytes.Buffer, err error) {
	err = jpeg.Encode(&buf, s.Img, &jpeg.Options{Quality: 90})
	return
}

func New(width int, height int) *Scene {
	return &Scene{
		Width:  width,
		Height: height,
		Img:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}
