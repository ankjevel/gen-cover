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

func parseFont(path string, size float64) font.Face {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	ttf_font, err := freetype.ParseFont(bytes)
	if err != nil {
		panic(err)
	}

	return truetype.NewFace(ttf_font, &truetype.Options{
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
