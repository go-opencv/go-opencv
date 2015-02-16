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

func TestGcvCalibrateCamera(t *testing.T) {
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

	camMat, rvec, tvec := GcvCalibrateCamera(objPts, imgPts, camMat)
	spew.Dump(camMat)
	spew.Dump(rvec)
	spew.Dump(tvec)
}
