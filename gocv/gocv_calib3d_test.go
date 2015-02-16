package gocv

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gonum/matrix/mat64"
)

func TestGcvInitCameraMatrix2D(t *testing.T) {
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
	spew.Dump(camMat)
}

// func TestGcvCalibrateCamera(t *testing.T) {
// 	objPts := NewGcvPoint3fVector(int64(4))
// 	objPts.Set(0, NewGcvPoint3f(0, 25, 0))
// 	objPts.Set(1, NewGcvPoint3f(0, -25, 0))
// 	objPts.Set(2, NewGcvPoint3f(-47, 25, 0))
// 	objPts.Set(3, NewGcvPoint3f(-47, -25, 0))

// 	imgPts := NewGcvPoint2fVector(int64(4))
// 	imgPts.Set(0, NewGcvPoint2f(1136.4140625, 1041.89208984))
// 	imgPts.Set(1, NewGcvPoint2f(1845.33190918, 671.39581299))
// 	imgPts.Set(2, NewGcvPoint2f(302.73373413, 634.79998779))
// 	imgPts.Set(3, NewGcvPoint2f(1051.46154785, 352.76107788))

// 	imgSize := NewGcvSize2i(1920, 1080)

// 	camMat := GcvInitCameraMatrix2D(objPts, imgPts)
// 	spew.Dump(camMat.GcvAtd(NewGcvSize2i(0, 0)))
// 	spew.Dump(camMat.GcvAtd(NewGcvSize2i(0, 1)))
// 	spew.Dump(camMat.GcvAtd(NewGcvSize2i(1, 1)))
// 	spew.Dump(camMat.GcvAtd(NewGcvSize2i(1, 2)))
// 	spew.Dump(camMat.GcvAtd(NewGcvSize2i(2, 2)))

// 	rvec := NewMat()
// 	tvec := NewMat()

// 	GcvCalibrateCamera(objPts, imgPts, imgSize, camMat, rvec, tvec)

// 	MatToMat64(camMat)
// }
