package main

import (
	"fmt"
	"os"
	//"path"
	//"runtime"

	"github.com/hybridgroup/go-opencv/opencv"
	//"../opencv" // can be used in forks, comment in real application
)

func main() {
	win := opencv.NewWindow("Go-OpenCV Webcam")
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("can not open camera")
	}
	defer cap.Release()

	win.CreateTrackbar("Thresh", 1, 100, func(pos int, param ...interface{}) {
		for {
			if cap.GrabFrame() {
				img := cap.RetrieveFrame(1)
				if img != nil {
					ProcessImage(img, win, pos)
				} else {
					fmt.Println("Image ins nil")
				}
			}

			if key := opencv.WaitKey(10); key == 27 {
				os.Exit(0)
			}
		}
	})
	opencv.WaitKey(0)
}

func ProcessImage(img *opencv.IplImage, win *opencv.Window, pos int) error {
	w := img.Width()
	h := img.Height()

	// Create the output image
	cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
	defer cedge.Release()

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	defer gray.Release()
	defer edge.Release()

	opencv.CvtColor(img, gray, opencv.CV_BGR2GRAY)

	opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
	opencv.Not(gray, edge)

	// Run the edge detector on grayscale
	opencv.Canny(gray, edge, float64(pos), float64(pos*3), 3)

	opencv.Zero(cedge)
	// copy edge points
	opencv.Copy(img, cedge, edge)

	win.ShowImage(cedge)
	return nil
}
