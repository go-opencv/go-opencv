//
// thresholds.go
//
// Sample showing the different types of thresholding available.
// It does both fixed and adaptive thresholding.
//
// kevinabrandon@gmail.com - 9-14-2016
//

package main

import (
	"fmt"
	"github.com/lazywei/go-opencv/opencv"
	"os"
)

func main() {
	filename := "../images/lena.jpg"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	fmt.Println("Press ESC or 'q' to quit")

	img := opencv.LoadImage(filename, opencv.CV_LOAD_IMAGE_GRAYSCALE)
	if img == nil {
		fmt.Printf("LoadImage fail: %s\n", filename)
		return
	}
	defer img.Release()

	win := opencv.NewWindow("Threshold Sample")
	defer win.Destroy()

	win.CreateTrackbar("ThresholdType", 0, 4, func(pos int) { ThresholdImage(img, win) })
	win.CreateTrackbar("Threshold", 128, 255, func(pos int) { ThresholdImage(img, win) })
	win.CreateTrackbar("MaxValue", 255, 255, func(pos int) { ThresholdImage(img, win) })
	win.CreateTrackbar("Fixed/Adaptive", 0, 1, func(pos int) { ThresholdImage(img, win) })
	win.CreateTrackbar("AdaptiveMethod", 0, 1, func(pos int) { ThresholdImage(img, win) })
	win.CreateTrackbar("AdaptiveBlockSize", 3, 100, func(pos int) { ThresholdImage(img, win) })

	ThresholdImage(img, win)

	for {
		key := opencv.WaitKey(10)

		if key == 1048603 || key == 1048689 {
			// esc or 'q' key pressed
			break
		}
	}
}

var (
	threshTypes = []string{
		"Binary",          // opencv.CV_THRESH_BINARY
		"Binary Inverse",  // opencv.CV_THRESH_BINARY_INV
		"Truncated",       // opencv.CV_THRESH_TRUNC
		"To Zero",         // opencv.CV_THRESH_TOZERO
		"To Zero Inverse", // opencv.CV_THRESH_TOZERO_INV
	}

	adaptiveMethods = []string{
		"Mean",     // opencv.CV_ADAPTIVE_THRESH_MEAN_C
		"Gaussian", // opencv.CV_ADAPTIVE_THRESH_GAUSSIAN_C
	}
)

func ThresholdImage(img *opencv.IplImage, win *opencv.Window) {

	threshImg := opencv.CreateImage(img.Width(), img.Height(), opencv.IPL_DEPTH_8U, 1)
	defer threshImg.Release()

	threshType, _ := win.GetTrackbarPos("ThresholdType")
	thresh, _ := win.GetTrackbarPos("Threshold")
	maxVal, _ := win.GetTrackbarPos("MaxValue")
	fixedOrAdaptive, _ := win.GetTrackbarPos("Fixed/Adaptive")
	adaptMethd, _ := win.GetTrackbarPos("AdaptiveMethod")
	blockSize, _ := win.GetTrackbarPos("AdaptiveBlockSize")

	if fixedOrAdaptive == 0 {
		// fixed threshold:
		fmt.Println("*************************")
		fmt.Printf("Fixed threshold: %d\n", thresh)
		fmt.Printf("maxValue: %d\n", maxVal)
		fmt.Printf("thresType: '%s'\n", threshTypes[threshType])

		opencv.Threshold(
			img,
			threshImg,
			float64(thresh),
			float64(maxVal),
			threshType,
		)
	} else {
		// check to make sure we have a valid threshold type. (Binary or Binary inverse only)
		if threshType > 1 {
			fmt.Println("*** Only threshold types of Binary or Binary Inverse are allowed for adaptive thresholds!")
			return
		}

		// blockSize must be odd and >= 3
		if blockSize < 3 {
			blockSize = 3
		}
		if blockSize%2 == 0 {
			blockSize++
		}
		fmt.Println("*************************")
		fmt.Printf("Adaptive threshold: %d\n", thresh)
		fmt.Printf("blockSize: %d\n", blockSize)
		fmt.Printf("maxValue: %d\n", maxVal)
		fmt.Printf("threshType: '%s'\n", threshTypes[threshType])
		fmt.Printf("adaptiveMethod: '%s'\n", adaptiveMethods[adaptMethd])
		// adaptive threshold:
		opencv.AdaptiveThreshold(
			img,
			threshImg,
			float64(maxVal),
			adaptMethd,
			threshType,
			blockSize,
			float64(thresh),
		)
	}

	win.ShowImage(threshImg)
}
