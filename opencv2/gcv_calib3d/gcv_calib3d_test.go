package gcv_calib3d

import "testing"

import "github.com/lazywei/go-opencv/opencv2/gcv_core"

// [[[  0.  25.   0.]
//   [  0. -25.   0.]
//   [-47.  25.   0.]
//   [-47. -25.   0.]]]
// [[[ 1136.4140625   1041.89208984]
//   [ 1845.33190918   671.39581299]
//   [  302.73373413   634.79998779]
//   [ 1051.46154785   352.76107788]]]
// (1920, 1080)
// [[  4.82812906e+03   0.00000000e+00   9.59500000e+02]
//  [  0.00000000e+00   4.82812906e+03   5.39500000e+02]
//  [  0.00000000e+00   0.00000000e+00   1.00000000e+00]]

func TestGcvInitCameraMatrix2D(t *testing.T) {
	objPts := gcv_core.NewGcvPoint3fVector(int64(4))
	objPts.Set(0, gcv_core.NewGcvPoint3f(
		float32(0), float32(25), float32(0)))
	objPts.Set(1, gcv_core.NewGcvPoint3f(
		float32(0), float32(-25), float32(0)))
	objPts.Set(2, gcv_core.NewGcvPoint3f(
		float32(-47), float32(25), float32(0)))
	objPts.Set(3, gcv_core.NewGcvPoint3f(
		float32(-47), float32(-25), float32(0)))

	imgPts := gcv_core.NewGcvPoint2fVector(int64(4))
	imgPts.Set(0, gcv_core.NewGcvPoint2f(
		float32(1136.4140625), float32(1041.89208984)))
	imgPts.Set(1, gcv_core.NewGcvPoint2f(
		float32(1845.33190918), float32(671.39581299)))
	imgPts.Set(2, gcv_core.NewGcvPoint2f(
		float32(302.73373413), float32(634.79998779)))
	imgPts.Set(3, gcv_core.NewGcvPoint2f(
		float32(1051.46154785), float32(352.76107788)))

	GcvInitCameraMatrix2D(objPts, imgPts)
}

func TestGcvCalibrateCamera(t *testing.T) {
	objPts := gcv_core.NewGcvPoint3fVector(int64(4))
	objPts.Set(0, gcv_core.NewGcvPoint3f(
		float32(0), float32(25), float32(0)))
	objPts.Set(1, gcv_core.NewGcvPoint3f(
		float32(0), float32(-25), float32(0)))
	objPts.Set(2, gcv_core.NewGcvPoint3f(
		float32(-47), float32(25), float32(0)))
	objPts.Set(3, gcv_core.NewGcvPoint3f(
		float32(-47), float32(-25), float32(0)))

	imgPts := gcv_core.NewGcvPoint2fVector(int64(4))
	imgPts.Set(0, gcv_core.NewGcvPoint2f(
		float32(1136.4140625), float32(1041.89208984)))
	imgPts.Set(1, gcv_core.NewGcvPoint2f(
		float32(1845.33190918), float32(671.39581299)))
	imgPts.Set(2, gcv_core.NewGcvPoint2f(
		float32(302.73373413), float32(634.79998779)))
	imgPts.Set(3, gcv_core.NewGcvPoint2f(
		float32(1051.46154785), float32(352.76107788)))

	imgSize := gcv_core.NewGcvSize2i(1920, 1080)

	camMat := GcvInitCameraMatrix2D(objPts, imgPts)

	GcvCalibrateCamera(objPts, imgPts, imgSize, camMat)
}
