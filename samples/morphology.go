//
// morphology.go
//
// Sample showing how the different morphological operations work.
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
	filename := "../images/j.png"
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

	win := opencv.NewWindow("Morphology Sample", opencv.CV_WINDOW_NORMAL)
	defer win.Destroy()

	callback := func(pos int) { ProcessImage(img, win) }

	win.CreateTrackbar("NErode", 0, 50, callback)
	win.CreateTrackbar("NDilate", 0, 50, callback)
	win.CreateTrackbar("NMorphEx", 0, 50, callback)
	win.CreateTrackbar("MorphExTypes", 0, 4, callback)
	win.CreateTrackbar("StructWidth", 3, 50, callback)
	win.CreateTrackbar("StructHeight", 3, 50, callback)
	win.CreateTrackbar("StructType", 0, 2, callback)

	ProcessImage(img, win)

	for {
		key := opencv.WaitKey(10)

		if key == 1048603 || key == 1048689 {
			// esc or 'q' key pressed
			break
		}
	}
}

type NamedConsts struct {
	Name string
	C    int
}

var (
	morphTypes = []NamedConsts{
		NamedConsts{"Open", opencv.CV_MORPH_OPEN},
		NamedConsts{"Close", opencv.CV_MORPH_CLOSE},
		NamedConsts{"Gradient", opencv.CV_MORPH_GRADIENT},
		NamedConsts{"Top Hat", opencv.CV_MORPH_TOPHAT},
		NamedConsts{"Black Hat", opencv.CV_MORPH_BLACKHAT},
	}
	structShapes = []NamedConsts{
		NamedConsts{"Rectangle", opencv.CV_MORPH_RECT},
		NamedConsts{"Ellipse", opencv.CV_MORPH_ELLIPSE},
		NamedConsts{"Cross", opencv.CV_MORPH_CROSS},
	}
)

func ProcessImage(img *opencv.IplImage, win *opencv.Window) {

	nErode, _ := win.GetTrackbarPos("NErode")
	nDilate, _ := win.GetTrackbarPos("NDilate")
	nMorphEx, _ := win.GetTrackbarPos("NMorphEx")
	morphIndx, _ := win.GetTrackbarPos("MorphExTypes")
	sWidth, _ := win.GetTrackbarPos("StructWidth")
	sHeight, _ := win.GetTrackbarPos("StructHeight")
	sIndx, _ := win.GetTrackbarPos("StructType")

	// create the structuring element:
	var element *opencv.IplConvKernel
	if sWidth != 0 && sHeight != 0 {
		element = opencv.CreateStructuringElement(
			sWidth,                // width
			sHeight,               // height
			sWidth/2,              // X anchor
			sHeight/2,             // Y anchor
			structShapes[sIndx].C, // shape constant
		)
		defer element.ReleaseElement()
	}

	fmt.Println("****************************************")
	fmt.Printf("NErode: %d\n", nErode)
	fmt.Printf("nDilate: %d\n", nDilate)
	fmt.Printf("nMorphEx: %d\n", nMorphEx)
	fmt.Printf("morphType: %v\n", morphTypes[morphIndx])
	if element == nil {
		fmt.Println("Structuring element: nil (default 3x3 rect)")
	} else {
		fmt.Printf("Structuring element: [%dx%d] %v\n", sWidth, sHeight, structShapes[sIndx])
	}

	outputImg := img.Clone()
	defer outputImg.Release()

	// first we will erode the image..
	if nErode > 0 {
		opencv.Erode(
			outputImg, // source
			outputImg, // destination (inplace is okay here)
			element,   // structuring element (nil means default 3x3)
			nErode,    // number of iterations
		)
	}

	// next we will dilate the image...
	if nDilate > 0 {
		opencv.Dilate(
			outputImg, // source
			outputImg, // destination (inplace is okay here)
			element,   // structuring element (nil means default 3x3)
			nDilate,   // number of iterations
		)
	}

	// last we'll do the MorphologyEx()
	if nMorphEx > 0 {
		tempImg := opencv.CreateImage(img.Width(), img.Height(), opencv.IPL_DEPTH_8U, 1)
		defer tempImg.Release()

		opencv.MorphologyEx(
			outputImg,               // source image
			outputImg,               // destination (inplace is okay)
			tempImg,                 // temporary image
			element,                 // structuring element (nil means default 3x3)
			morphTypes[morphIndx].C, // morphological operation constant
			nMorphEx,                // number of iterations
		)
	}

	win.ShowImage(outputImg)
}
