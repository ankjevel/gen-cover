package fonts

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

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

func parseFont(path string, size float64) font.Face {
	dat, err := os.ReadFile(path)
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
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		log.Println("could not get current filename")
		panic(0)
	}

	src_dir := filepath.Dir(filename) + "/src"

	FaceSub = parseFont(src_dir+"/Lato-Light.ttf", FaceSubSize)
	FaceMain = parseFont(src_dir+"/Lato-Black.ttf", FaceMainSize)
}
