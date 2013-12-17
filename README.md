#Go OpenCV binding

## Why fork?
Original project (https://code.google.com/p/go-opencv) looks abandoned. Therefore, I decide to fork it and host it on Github, so that others can help to maintain this package.

## Install

### Linux
- install Go1 and OpenCV
- `go get github.com/lazywei/go-opencv`
- `cd ${GoOpenCVRoot}/samples && go run hellocv.go`

### Windows
- install Go1 and MinGW
- install OpenCV-2.4.x to MinGW dir
  - libopencv*.dll --> ${MinGWRoot}\bin
  - libopencv*.lib --> ${MinGWRoot}\lib
  - include\opencv --> ${MinGWRoot}\include\opencv
  - include\opencv2 --> ${MinGWRoot}\include\opencv2
- go get code.google.com/p/go-opencv/trunk/opencv
- cd ${GoOpenCVRoot}/trunk/samples && go run hellocv.go

## Example

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
You can find more samples at: https://github.com/lazywei/go-opencv/tree/master/samples

## Contribute

- Fork this repo
- Create new feature branch, `git checkout -b your-feature-branch`
- Commit your change and push it to your repo `git commit -m 'new feature'; git push origin your-feature-branch`
- Open pull request!

## TODOs
- More details doc
- Implement more bindings
