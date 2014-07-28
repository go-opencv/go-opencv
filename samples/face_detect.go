package main

import (
	"github.com/lazywei/go-opencv/opencv"
	"path"
	"runtime"
)

func main() {
	_, currentfile, _, _ := runtime.Caller(0)
	image := opencv.LoadImage(path.Join(path.Dir(currentfile), "../images/lena.jpg"))

	cascade := opencv.LoadHaarClassifierCascade(path.Join(path.Dir(currentfile), "haarcascade_frontalface_alt.xml"))
	faces := cascade.DetectObjects(image)

	for _, value := range faces {
		opencv.Rectangle(image,
			opencv.Point{value.X() + value.Width(), value.Y()},
			opencv.Point{value.X(), value.Y() + value.Height()},
			opencv.ScalarAll(255.0), 1, 1, 0)
	}

	win := opencv.NewWindow("Face Detection")
	win.ShowImage(image)
	opencv.WaitKey(0)
}
