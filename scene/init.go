package scene

import (
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		log.Println("could not get current filename")
		panic(0)
	}

	static_dir := filepath.Dir(filename) + "/../static"

	noise = decode(static_dir + "/noise.png")

}
