package main

import (
	"os"
	"fmt"
	"path"
	"runtime"
	"go-opencv/opencv"
	//opencv "github.com/tbolsh/go-opencv/opencv"
)
// see http://docs.opencv.org/2.4/doc/tutorials/imgproc/imgtrans/hough_lines/hough_lines.html
func main() {
	fmt.Println( "\nThis program demonstrates line finding with the Hough transform.")
    fmt.Println("Usage:")
    fmt.Println("./lines_detect <image_name>, Default is Hough_Lines_Tutorial_Original_Image.jpg");
	fmt.Println("To exit please press ESCAPE when the image window is in focus.");


	 _, currentfile, _, _ := runtime.Caller(0)
	filename := path.Join(path.Dir(currentfile), "../images/Hough_Lines_Tutorial_Original_Image.jpg")
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	image := opencv.LoadImage(filename)
	if image == nil {
		panic("LoadImage fail")
	}
	defer image.Release()
	
	w, h := image.Width(), image.Height()
	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	defer gray.Release()
	opencv.CvtColor(image, gray, opencv.CV_BGR2GRAY)

	// find edges
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	defer edge.Release()
	edge_thresh := 50.0
	opencv.Not(gray, edge)
	opencv.Canny(gray, edge, edge_thresh, edge_thresh*4.0, 3)

	cimage := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
	defer cimage.Release()
	
	win := opencv.NewWindow("Lines: Probabilistic Hough Line transform demo")
	defer win.Destroy()

	color := opencv.NewScalar(0 ,0, 255, 1)
	win.CreateTrackbar("Thresh", 1, 100, func(pos int, param ...interface{}) {
		thresh := pos
	
		opencv.Zero(cimage)
		opencv.Copy(image, cimage, nil)
		// find lines
		lines := opencv.HoughLinesP(edge, 1.0, opencv.CV_PI/180, 50+thresh, 50.0, 10.0)
		for _,l := range lines {
			opencv.Line(cimage, l.P1.Point(), l.P2.Point(), color, 3, 0, 0)
		}
		win.ShowImage(cimage)
	})
	win.ShowImage(image)
	
	for {
		key := opencv.WaitKey(20)
		if key == 27 {
			os.Exit(0)
		}
	}

	os.Exit(0)
}