Go OpenCV binding
==================

A Golang binding for [OpenCV](http://opencv.org/).

[**DISCLAIMER**](https://github.com/lazywei/go-opencv#disclaimer)

## Install

### Linux & Mac OS X

Install Go and OpenCV, you might want to install both of them via `apt-get` or `homebrew`.

```
go get github.com/lazywei/go-opencv
cd ${GoOpenCVRoot}/samples && go run hellocv.go
```

### Windows

- Install Go and MinGw
- install OpenCV-2.4.x to MinGW dir

```
# libopencv*.dll --> ${MinGWRoot}\bin
# libopencv*.lib --> ${MinGWRoot}\lib
# include\opencv --> ${MinGWRoot}\include\opencv
# include\opencv2 --> ${MinGWRoot}\include\opencv2

go get code.google.com/p/go-opencv/trunk/opencv
cd ${GoOpenCVRoot}/trunk/samples && go run hellocv.go
```

## Usage

Currently there are no too many readily instruction for usage. At this point, you can always refers to OpenCV's documentation. I'll try to keep all the bindings have the same signature as in OpenCV's C interface. However, please do note that sometimes the signature might slightly differ from the C interface due to Golang's type declaration conventions, for example:

```
# The original signature in C interface.
void cvInitIntrinsicParams2D(
const CvMat* object_points, const CvMat* image_points,
const CvMat* npoints, CvSize image_size, CvMat* camera_matrix,
double aspect_ratio=1. )

# We might put all *Mat types together, however
func InitIntrinsicParams2D(objectPoints, imagePoints, nPoints, cameraMatrix *Mat ...

# Or we might use "explicitly return" instead of C-style's pointer
func InitIntrinsicParams2D(objectPoints, imagePoints, nPoints, *Mat ...) (cameraMatrix *Mat)
```

## TODOs

- [ ] Better documents
- [ ] Split the big package into sub-packages corresponding to the modules described in [OpenCV API Reference](http://docs.opencv.org/modules/core/doc/intro.html)
- [ ] Clean up the codes with better coding style

## Example

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

-------------------

## Disclaimer

This is a fork of [chai's go-opencv](https://github.com/chai2010/opencv). At the time of the fork (Dec 9, 2013) the original project was inactive, and hence I decide to host a fork on Github so people can contribute to this project easily. However, now it seems to be active again starting from Aug 25, 2014. Efforts to merge the two projects are very welcome.

