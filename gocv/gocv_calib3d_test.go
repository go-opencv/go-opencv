package gocv

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/stretchr/testify/assert"
)

func TestGcvInitCameraMatrix2D(t *testing.T) {
	objPts := mat64.NewDense(4, 3, []float64{
		0, 25, 0,
		0, -25, 0,
		-47, 25, 0,
		-47, -25, 0,
	})

	imgPts := mat64.NewDense(4, 2, []float64{
		1136.4140625, 1041.89208984,
		1845.33190918, 671.39581299,
		302.73373413, 634.79998779,
		1051.46154785, 352.76107788,
	})

	camMat := GcvInitCameraMatrix2D(objPts, imgPts)
	assert.Equal(t, camMat.Row(nil, 0), []float64{4828.129063751587, 0, 959.5})
	assert.Equal(t, camMat.Row(nil, 1), []float64{0, 4828.129063751587, 539.5})
	assert.Equal(t, camMat.Row(nil, 2), []float64{0, 0, 1})
}

func TestGcvCalibrateCamera(t *testing.T) {
	objPts := mat64.NewDense(4, 3, []float64{
		0, 25, 0,
		0, -25, 0,
		-47, 25, 0,
		-47, -25, 0,
	})

	imgPts := mat64.NewDense(4, 2, []float64{
		1136.4140625, 1041.89208984,
		1845.33190918, 671.39581299,
		302.73373413, 634.79998779,
		1051.46154785, 352.76107788,
	})

	camMat := GcvInitCameraMatrix2D(objPts, imgPts)

	camMat, rvec, tvec := GcvCalibrateCamera(objPts, imgPts, camMat)

	assert.Equal(t, camMat.Row(nil, 0), []float64{5.47022369e+03, 0.00000000e+00, 9.59500000e+02})
	assert.Equal(t, camMat.Row(nil, 1), []float64{0.00000000e+00, 5.47022369e+03, 5.39500000e+02})
	assert.Equal(t, camMat.Row(nil, 2), []float64{0.00000000e+00, 0.00000000e+00, 1.00000000e+00})

	assert.Equal(t, rvec.Col(nil, 0), []float64{-0.99458984, 0.54674764, -2.69721055})
	assert.Equal(t, tvec.Col(nil, 0), []float64{-23.25417757, -12.6155423, -227.64212085})
}

func TestGcvRodrigues(t *testing.T) {
	rvec := mat64.NewDense(3, 1, []float64{
		-0.99458984,
		0.54674764,
		-2.69721055,
	})
	rmat := GcvRodrigues(rvec)

	assert.InDelta(t, rmat.At(0, 0), -0.74853587, 1e-6)
	assert.InDelta(t, rmat.At(0, 1), 0.07139127, 1e-6)
	assert.InDelta(t, rmat.At(0, 2), 0.65923997, 1e-6)

	assert.InDelta(t, rmat.At(1, 0), -0.32247419, 1e-6)
	assert.InDelta(t, rmat.At(1, 1), -0.90789575, 1e-6)
	assert.InDelta(t, rmat.At(1, 2), -0.26783521, 1e-6)

	assert.InDelta(t, rmat.At(2, 0), 0.57940008, 1e-6)
	assert.InDelta(t, rmat.At(2, 1), -0.41307214, 1e-6)
	assert.InDelta(t, rmat.At(2, 2), 0.70261437, 1e-6)
}
