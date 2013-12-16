#Go OpenCV binding

## Why fork?
Original project (https://code.google.com/p/go-opencv) looks abandoned. Therefore, I decide to fork it and host it on Github, so that others can help to maintain this package.

## TODOs
- More details doc
- Implement more bindings

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
