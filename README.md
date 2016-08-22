Go OpenCV binding
==================

[![Join the chat at https://gitter.im/lazywei/go-opencv](https://badges.gitter.im/lazywei/go-opencv.svg)](https://gitter.im/lazywei/go-opencv?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

A Golang binding for [OpenCV](http://opencv.org/).

OpenCV 1.x C API bindings through CGO, and OpenCV 2+ C++ API ([`GoCV`](gocv/)) through SWIG.

-------------------

## Disclaimer

This is a fork of [chai's go-opencv](https://github.com/chai2010/opencv), which has only OpenCV1 support through CGO, and all credits for OpenCV1 wrapper (except files in `gocv/` folder) should mainly go to Chai. At the time of the fork (Dec 9, 2013) the original project was inactive and was hosted on Google Code, which was a little inconvenient for community contribution. Hence, I decided to host a fork on Github so people can contribute to this project easily. Since then, some patches were added by community, and some experimental OpenCV 2 wrappers were added as well. That means this fork went on a little bit divergent way comparing to the origin project. However, now the origin project seems to be active again and be moved to GitHub starting from Aug 25, 2014. Efforts to merge the two projects are very welcome.

-------------------

## Install

### Linux & Mac OS X

Install Go and OpenCV, you might want to install both of them via `apt-get` or `homebrew`.

```
go get github.com/lazywei/go-opencv
cd $GOPATH/src/github.com/lazywei/go-opencv/samples
go run hellocv.go
```

### Windows

- Install Go and MinGw
- install OpenCV-2.4.x to MinGW dir

```
# libopencv*.dll --> ${MinGWRoot}\bin
# libopencv*.lib --> ${MinGWRoot}\lib
# include\opencv --> ${MinGWRoot}\include\opencv
# include\opencv2 --> ${MinGWRoot}\include\opencv2

go get github.com/lazywei/go-opencv
cd ${GoOpenCVRoot}/trunk/samples && go run hellocv.go
```

## [WIP] OpenCV2 (GoCV)

After OpenCV 2.x+, the core team no longer develop and maintain C API. Therefore, CGO will not be used in CV2 binding. Instead, we are using SWIG for wrapping. The support for OpenCV2 is currently under development, and whole code will be placed under `gocv` package.

If you want to use CV2's API, please refer to the code under `gocv/` directory. There is no too many documents for CV2 wrapper yet, but you can still find the example usages in `*_test.go`.

Please also note that the basic data structures in OpenCV (e.g., `cv::Mat`, `cv::Point3f`) are wrapped partially for now. For more detail on how to use these types, please refer to [GoCV's README](gocv/README.md).

*Requirement*: we will build the wrappers based on [mat64](https://godoc.org/github.com/gonum/matrix/mat64), given it is much easier to manipulate the underlaying data. In most case, it is not necessary to access the original CV data, e.g., `cv::Mat` can be converted from/to `*mat64.Dense`.

## Example

### OpenCV2's initCameraMatrix2D

```go
package main

import . "github.com/lazywei/go-opencv/gocv"
import "github.com/gonum/matrix/mat64"

func main() {

	objPts := mat64.NewDense(4, 3, []float64{
		0, 25, 0,
		0, -25, 0,
		-47, 25, 0,
		-47, -25, 0})

	imgPts := mat64.NewDense(4, 2, []float64{
		1136.4140625, 1041.89208984,
		1845.33190918, 671.39581299,
		302.73373413, 634.79998779,
		1051.46154785, 352.76107788})

	camMat := GcvInitCameraMatrix2D(objPts, imgPts)
	fmt.Println(camMat)
}
```


### Resizing

```go
package main

import opencv "github.com/lazywei/go-opencv/opencv"

func main() {
	filename := "bert.jpg"
	srcImg := opencv.LoadImage(filename)
	if srcImg == nil {
		panic("Loading Image failed")
	}
	defer srcImg.Release()
	resized1 := opencv.Resize(srcImg, 400, 0, 0)
	resized2 := opencv.Resize(srcImg, 300, 500, 0)
	resized3 := opencv.Resize(srcImg, 300, 500, 2)
	opencv.SaveImage("resized1.jpg", resized1, 0)
	opencv.SaveImage("resized2.jpg", resized2, 0)
	opencv.SaveImage("resized3.jpg", resized3, 0)
}
```

### Webcam

Yet another cool example is created by @saratovsource which demos how to use webcam:

```
cd samples
go run webcam.go
```

### More

You can find more samples at: https://github.com/lazywei/go-opencv/tree/master/samples

## How to contribute

- Fork this repo
- Clone the main repo, and add your fork as a remote

  ```
  git clone https://github.com/lazywei/go-opencv.git
  cd go-opencv
  git remote rename origin upstream
  git remote add origin https://github.com/your_github_account/go-opencv.git
  ```

- Create new feature branch

  ```
  git checkout -b your-feature-branch
  ```

- Commit your change and push it to your repo 

  ```
  git commit -m 'new feature'
  git push origin your-feature-branch
  ```

- Open a pull request!

