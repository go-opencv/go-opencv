// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/lazywei/go-opencv/opencv"
	//"../opencv" // can be used in forks, comment in real application
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

	w := image.Width()
	h := image.Height()

	// Create the output image
	cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
	defer cedge.Release()

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	defer gray.Release()
	defer edge.Release()

	opencv.CvtColor(image, gray, opencv.CV_BGR2GRAY)

	win := opencv.NewWindow("Edge")
	defer win.Destroy()

	win.SetMouseCallback(func(event, x, y, flags int, param ...interface{}) {
		fmt.Printf("event = %d, x = %d, y = %d, flags = %d\n",
			event, x, y, flags,
		)
	})

	win.CreateTrackbar("Thresh", 1, 100, func(pos int, param ...interface{}) {
		edge_thresh := pos

		opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
		opencv.Not(gray, edge)

		// Run the edge detector on grayscale
		opencv.Canny(gray, edge, float64(edge_thresh), float64(edge_thresh*3), 3)

		opencv.Zero(cedge)
		// copy edge points
		opencv.Copy(image, cedge, edge)

		win.ShowImage(cedge)

		fmt.Printf("pos = %d\n", pos)
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
