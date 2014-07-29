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
	filename := path.Join(path.Dir(currentfile), "../images/fruits.jpg")
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	img0 := opencv.LoadImage(filename)
	if img0 == nil {
		panic("LoadImage fail")
	}
	defer img0.Release()

	fmt.Print("Hot keys: \n",
		"\tESC - quit the program\n",
		"\tr - restore the original image\n",
		"\ti or ENTER - run inpainting algorithm\n",
		"\t\t(before running it, paint something on the image)\n",
	)

	img := img0.Clone()
	inpainted := img0.Clone()
	inpaint_mask := opencv.CreateImage(img0.Width(), img0.Height(), 8, 1)

	opencv.Zero(inpaint_mask)
	//opencv.Zero( inpainted )

	win := opencv.NewWindow("image")
	defer win.Destroy()

	prev_pt := opencv.Point{-1, -1}
	win.SetMouseCallback(func(event, x, y, flags int, param ...interface{}) {
		if img == nil {
			os.Exit(0)
		}

		if event == opencv.CV_EVENT_LBUTTONUP ||
			(flags&opencv.CV_EVENT_FLAG_LBUTTON) == 0 {
			prev_pt = opencv.Point{-1, -1}
		} else if event == opencv.CV_EVENT_LBUTTONDOWN {
			prev_pt = opencv.Point{x, y}
		} else if event == opencv.CV_EVENT_MOUSEMOVE &&
			(flags&opencv.CV_EVENT_FLAG_LBUTTON) != 0 {
			pt := opencv.Point{x, y}
			if prev_pt.X < 0 {
				prev_pt = pt
			}

			rgb := opencv.ScalarAll(255.0)
			opencv.Line(inpaint_mask, prev_pt, pt, rgb, 5, 8, 0)
			opencv.Line(img, prev_pt, pt, rgb, 5, 8, 0)
			prev_pt = pt

			win.ShowImage(img)
		}
	})
	win.ShowImage(img)
	opencv.WaitKey(0)

	win2 := opencv.NewWindow("inpainted image")
	defer win2.Destroy()
	win2.ShowImage(inpainted)

	for {
		key := opencv.WaitKey(20)
		if key == 27 {
			os.Exit(0)
		} else if key == 'r' {
			opencv.Zero(inpaint_mask)
			opencv.Copy(img0, img, nil)
			win.ShowImage(img)
		} else if key == 'i' || key == '\n' {
			opencv.Inpaint(img, inpaint_mask, inpainted, 3,
				opencv.CV_INPAINT_TELEA,
			)
			win2.ShowImage(inpainted)
		}
	}
	os.Exit(0)
}
