package main

import (
	"fmt"
	"os"
	"path"

	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	win := opencv.NewWindow("Go-OpenCV Webcam Face Detection")
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("cannot open camera")
	}
	defer cap.Release()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cascade := opencv.LoadHaarClassifierCascade(path.Join(cwd, "haarcascade_frontalface_alt.xml"))

	fmt.Println("Press ESC to quit")
	for {
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)
			if img != nil {
				faces := cascade.DetectObjects(img)
				for _, value := range faces {
					opencv.Circle(img,
						opencv.Point{
							value.X() + (value.Width() / 2),
							value.Y() + (value.Height() / 2),
						},
						value.Width()/2,
						opencv.ScalarAll(255.0), 1, 1, 0)
				}

				win.ShowImage(img)
			} else {
				fmt.Println("nil image")
			}
		}
		key := opencv.WaitKey(1)

		if key == 27 {
			os.Exit(0)
		}
	}
}
