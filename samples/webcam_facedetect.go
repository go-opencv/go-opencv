package main

import (
	"fmt"
	"os"
	"path"

	"../opencv"
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
					upLeftPoint := opencv.Point{X: values.X(), Y: values.Y()}
					downRightPoint := opencv.Point{
						X: values.X() + values.Width(),
						Y: values.Y() + values.Height()
					}
					bboxColor := opencv.NewScalar(0, 255.0, 0)
					opencv.Rectangle(img, upLeftPoint, downRightPoint, bboxColor, 3, 0, 0)
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
