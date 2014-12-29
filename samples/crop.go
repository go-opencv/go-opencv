package main

import (
	"os"
	"path"
	"runtime"

	opencv "github.com/lazywei/go-opencv/opencv"
)

func main() {
	_, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/lena.jpg")
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	image := opencv.LoadImage(filename)
	if image == nil {
		panic("LoadImage fail")
	}
	defer image.Release()

	crop := opencv.Crop(image, 0, 0, 50, 50)
	opencv.SaveImage("/tmp/crop.jpg", crop, 0)
	crop.Release()

	os.Exit(0)
}
