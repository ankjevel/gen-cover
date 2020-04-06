package scene

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/Iteam1337/gen-cover/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Scene struct {
	Width, Height int
	Img           *image.RGBA
}

var (
	wejay image.Image
	noise image.Image
)

func decode(path string) (img image.Image) {
	r, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	img, err = png.Decode(r)
	if err != nil {
		panic(err)
	}
	return
}

func init() {
	wejay = decode("./static/wejay-bl-100.png")
	noise = decode("./static/noise.png")
}

func (s *Scene) AddLabel(x, y int, label string) {
	fontsize := float64(fonts.FaceMainSize)
	dotX := fixed.Int26_6(fontsize * 1.4 * 64)
	size := s.Img.Bounds().Size()
	drawer := func(y int, label string, face font.Face) {
		drawer := &font.Drawer{
			Dst:  s.Img,
			Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
			Face: face,
			Dot: fixed.Point26_6{
				X: dotX,
				Y: fixed.Int26_6(y),
			},
		}
		drawer.DrawString(label)
	}

	sy := float64(size.Y)
	drawer(int(sy-(fontsize*6.6))*64, "wejay", fonts.FaceSub)
	drawer(int(sy-(fontsize*4.8))*64, label, fonts.FaceMain)
}

func (s *Scene) EachPixel(offset int, colorFunction func(int, int) color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.Img.Set(x, y, colorFunction(x, y))
		}
	}

	bounds := s.Img.Bounds()

	wb := wejay.Bounds()
	wy := fonts.FaceMainSize / 2
	wrect := image.Rect(0, wy, wb.Dx(), wy+wb.Dy())

	m := image.NewRGBA(bounds)
	draw.Draw(m, bounds, s.Img, image.Point{0, 0}, draw.Src)
	draw.Draw(m, bounds, noise, image.Point{0, 0}, draw.Over)
	draw.Draw(m, wrect, wejay, image.Point{0, 0}, draw.Over)

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
