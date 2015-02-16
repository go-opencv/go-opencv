package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
import "C"
import "github.com/gonum/matrix/mat64"

// GcvInitCameraMatrix2D takes one N-by-3 matrix and one
// N-by-2 Matrix as input.
// Each row in the input matrix represents a point in real
// world (objPts) or in image (imgPts).
// Return: the camera matrix.
func GcvInitCameraMatrix2D(objPts, imgPts *mat64.Dense) (camMat *mat64.Dense) {
	nObjPts, objCol := objPts.Dims()
	nImgPts, imgCol := imgPts.Dims()

	if objCol != 3 || imgCol != 2 || nObjPts != nImgPts {
		panic("Invalid dimensions for objPts and imgPts")
	}

	objPtsVec := NewGcvPoint3f32Vector(int64(nObjPts))
	imgPtsVec := NewGcvPoint2f32Vector(int64(nObjPts))

	for i := 0; i < nObjPts; i++ {
		objPtsVec.Set(i, NewGcvPoint3f32(
			objPts.At(i, 0), objPts.At(i, 1), objPts.At(i, 2)))
	}

	for i := 0; i < nObjPts; i++ {
		imgPtsVec.Set(i, NewGcvPoint2f32(
			imgPts.At(i, 0), imgPts.At(i, 1)))
	}

	camMat = GcvMatToMat64(GcvInitCameraMatrix2D_(objPtsVec, imgPtsVec))
	return camMat
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
