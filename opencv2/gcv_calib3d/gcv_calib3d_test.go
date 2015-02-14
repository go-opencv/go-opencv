package gcv_calib3d

import "testing"

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
	objPts := NewGcvPoint3fVector(int64(4))
	objPts.Set(0, GetPoint3f(0, 25, 0))
	objPts.Set(1, GetPoint3f(0, -25, 0))
	objPts.Set(2, GetPoint3f(-47, 25, 0))
	objPts.Set(3, GetPoint3f(-47, -25, 0))

	imgPts := NewGcvPoint2fVector(int64(4))
	imgPts.Set(0, GetPoint2f(1136.4140625, 1041.89208984))
	imgPts.Set(1, GetPoint2f(1845.33190918, 671.39581299))
	imgPts.Set(2, GetPoint2f(302.73373413, 634.79998779))
	imgPts.Set(3, GetPoint2f(1051.46154785, 352.76107788))

	GcvInitCameraMatrix2D(objPts, imgPts)
}
