package gocv

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewGcvPoint3f(t *testing.T) {
	pt := NewGcvPoint3f(3, 1, 2)
	spew.Dump(pt)
}

func TestNewGcvPoint2f(t *testing.T) {
	pt := NewGcvPoint2f(3, 1)
	spew.Dump(pt)
}

func TestNewGcvSize2d(t *testing.T) {
	size := NewGcvSize2d(3, 1)
	spew.Dump(size)
}

func TestMat(t *testing.T) {
	mat := NewMat()
	mat2 := NewMat(mat)
	spew.Dump(mat2)
}
func TestGcvInitCameraMatrix2D(t *testing.T) {
	objPts := NewGcvPoint3fVector(int64(4))
	objPts.Set(0, NewGcvPoint3f(0, 25, 0))
	objPts.Set(1, NewGcvPoint3f(0, -25, 0))
	objPts.Set(2, NewGcvPoint3f(-47, 25, 0))
	objPts.Set(3, NewGcvPoint3f(-47, -25, 0))

	imgPts := NewGcvPoint2fVector(int64(4))
	imgPts.Set(0, NewGcvPoint2f(1136.4140625, 1041.89208984))
	imgPts.Set(1, NewGcvPoint2f(1845.33190918, 671.39581299))
	imgPts.Set(2, NewGcvPoint2f(302.73373413, 634.79998779))
	imgPts.Set(3, NewGcvPoint2f(1051.46154785, 352.76107788))

	camMat := GcvInitCameraMatrix2D(objPts, imgPts)
	spew.Dump(camMat.GcvAtd(NewGcvSize2i(0, 0)))
	spew.Dump(camMat.GcvAtd(NewGcvSize2i(0, 1)))
	spew.Dump(camMat.GcvAtd(NewGcvSize2i(1, 1)))
	spew.Dump(camMat.GcvAtd(NewGcvSize2i(1, 2)))
	spew.Dump(camMat.GcvAtd(NewGcvSize2i(2, 2)))
}

func TestGcvCalibrateCamera(t *testing.T) {
	objPts := NewGcvPoint3fVector(int64(4))
	objPts.Set(0, NewGcvPoint3f(0, 25, 0))
	objPts.Set(1, NewGcvPoint3f(0, -25, 0))
	objPts.Set(2, NewGcvPoint3f(-47, 25, 0))
	objPts.Set(3, NewGcvPoint3f(-47, -25, 0))

	imgPts := NewGcvPoint2fVector(int64(4))
	imgPts.Set(0, NewGcvPoint2f(1136.4140625, 1041.89208984))
	imgPts.Set(1, NewGcvPoint2f(1845.33190918, 671.39581299))
	imgPts.Set(2, NewGcvPoint2f(302.73373413, 634.79998779))
	imgPts.Set(3, NewGcvPoint2f(1051.46154785, 352.76107788))

	imgSize := NewGcvSize2i(1920, 1080)

	camMat := GcvInitCameraMatrix2D(objPts, imgPts)

	GcvCalibrateCamera(objPts, imgPts, imgSize, camMat)
}
