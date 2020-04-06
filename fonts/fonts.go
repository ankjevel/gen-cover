package fonts

import (
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	FaceSubSize  = 12
	FaceMainSize = 23
)

var (
	FaceSub  font.Face
	FaceMain font.Face
)

func parsePath(path string, size float64) font.Face {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	ttfFont, err := freetype.ParseFont(dat)
	if err != nil {
		panic(err)
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size: size,
		DPI:  118,
	})
}

func init() {
	FaceSub = parsePath("./fonts/src/Lato-Regular.ttf", FaceSubSize)
	FaceMain = parsePath("./fonts/src/Lato-Black.ttf", FaceMainSize)
}
