// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"runtime"

	// "github.com/lazywei/go-opencv/opencv"
	"../opencv" // can be used in forks, comment in real application
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

	mean, stdDev := image.MeanStdDev()
	fmt.Println(mean, stdDev)

	width := image.Width()
	height := image.Height()

	// Calculated mean and standard deviation
	for i := 0; i < 3; i++ {

		total := 0.0
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				total += image.Get2D(x, y).Val()[i]
			}
		}

		average := total / float64(width*height)

		total = 0.0
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				deviation := image.Get2D(x, y).Val()[i] - average
				total += math.Pow(deviation, 2)
			}
		}

		variance := total / float64(width*height)
		stdDevCalculated := math.Sqrt(variance)

		fmt.Println(stdDevCalculated-stdDev.Val()[i] < 0.0000000001)
		fmt.Println(stdDevCalculated, average)
	}

}
