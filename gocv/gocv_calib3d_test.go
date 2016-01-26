package gocv

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/stretchr/testify/assert"
)

const (
	DELTA float64 = 1e-5
)

func TestGcvInitCameraMatrix2D(t *testing.T) {
	objPts := mat64.NewDense(10, 3, []float64{
		-1.482676, -1.419348, 1.166475,
		-0.043819, -0.729445, 1.212821,
		0.960825, 1.147328, 0.485541,
		1.738245, 0.597865, 1.026016,
		-0.430206, -1.281281, 0.870726,
		-1.627323, -2.203264, -0.381758,
		0.166347, -0.571246, 0.428893,
		0.376266, 0.213996, -0.299131,
		-0.226950, 0.942377, -0.899869,
		-1.148912, 0.093725, 0.634745,
	})
	objPts.Clone(objPts.T())

	imgPts := mat64.NewDense(10, 2, []float64{
		-0.384281, -0.299055,
		0.361833, 0.087737,
		1.370253, 1.753933,
		1.421390, 0.853312,
		0.107177, -0.443076,
		3.773328, 5.437829,
		0.624914, -0.280949,
		-0.825577, -0.245594,
		0.631444, -0.340257,
		-0.647580, 0.502113,
	})
	imgPts.Clone(imgPts.T())

	camMat := GcvInitCameraMatrix2D(objPts, imgPts, [2]int{1920, 1080}, 1)
	assert.InDeltaSlice(t, []float64{1.47219772e+03, 0.00000000e+00, 9.59500000e+02},
		mat64.Row(nil, 0, camMat), DELTA)
	assert.InDeltaSlice(t, []float64{0.00000000e+00, 1.47219772e+03, 5.39500000e+02},
		mat64.Row(nil, 1, camMat), DELTA)
	assert.InDeltaSlice(t, []float64{0.00000000e+00, 0.00000000e+00, 1.00000000e+00},
		mat64.Row(nil, 2, camMat), DELTA)
}

func TestGcvCalibrateCamera(t *testing.T) {
	objPts := mat64.NewDense(10, 3, []float64{
		-1.482676, -1.419348, 1.166475,
		-0.043819, -0.729445, 1.212821,
		0.960825, 1.147328, 0.485541,
		1.738245, 0.597865, 1.026016,
		-0.430206, -1.281281, 0.870726,
		-1.627323, -2.203264, -0.381758,
		0.166347, -0.571246, 0.428893,
		0.376266, 0.213996, -0.299131,
		-0.226950, 0.942377, -0.899869,
		-1.148912, 0.093725, 0.634745,
	})
	objPts.Clone(objPts.T())

	imgPts := mat64.NewDense(10, 2, []float64{
		-0.384281, -0.299055,
		0.361833, 0.087737,
		1.370253, 1.753933,
		1.421390, 0.853312,
		0.107177, -0.443076,
		3.773328, 5.437829,
		0.624914, -0.280949,
		-0.825577, -0.245594,
		0.631444, -0.340257,
		-0.647580, 0.502113,
	})
	imgPts.Clone(imgPts.T())

	camMat := GcvInitCameraMatrix2D(objPts, imgPts, [2]int{1920, 1080}, 1)

	distCoeffs := mat64.NewDense(5, 1, []float64{0, 0, 0, 0, 0})

	camMat, rvec, tvec := GcvCalibrateCamera(
		objPts, imgPts, camMat, distCoeffs, [2]int{1920, 1080}, 14575)

	assert.InDeltaSlice(t, []float64{-46.15296606, 0., 959.5}, mat64.Row(nil, 0, camMat), DELTA)
	assert.InDeltaSlice(t, []float64{0., -46.15296606, 539.5}, mat64.Row(nil, 1, camMat), DELTA)
	assert.InDeltaSlice(t, []float64{0., 0., 1.}, mat64.Row(nil, 2, camMat), DELTA)

	assert.InDeltaSlice(t, []float64{-0.98405029, -0.93443411, -0.26304667}, mat64.Col(nil, 0, rvec), DELTA)
	assert.InDeltaSlice(t, []float64{0.6804739, 0.47530207, -0.04833094}, mat64.Col(nil, 0, tvec), DELTA)
}

func TestGcvRodrigues(t *testing.T) {
	rvec := mat64.NewDense(3, 1, []float64{
		-0.98405029,
		-0.93443411,
		-0.26304667,
	})
	rmat := GcvRodrigues(rvec)

	assert.InDeltaSlice(t, []float64{0.59922526, 0.57799222, -0.55394411}, mat64.Row(nil, 0, rmat), DELTA)
	assert.InDeltaSlice(t, []float64{0.20413818, 0.558743, 0.80382452}, mat64.Row(nil, 1, rmat), DELTA)
	assert.InDeltaSlice(t, []float64{0.77411672, -0.5947531, 0.21682264}, mat64.Row(nil, 2, rmat), DELTA)
}
