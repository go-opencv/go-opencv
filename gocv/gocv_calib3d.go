package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
// #cgo linux  pkg-config: opencv
import "C"
import "github.com/gonum/matrix/mat64"

// GcvInitCameraMatrix2D takes one 3-by-N matrix and one 2-by-N Matrix as input.
// Each column in the input matrix represents a point in real world (objPts) or
// in image (imgPts).
// Return: the camera matrix.
func GcvInitCameraMatrix2D(objPts, imgPts *mat64.Dense, dims [2]int,
	aspectRatio float64) (camMat *mat64.Dense) {

	objDim, nObjPts := objPts.Dims()
	imgDim, nImgPts := imgPts.Dims()

	if objDim != 3 || imgDim != 2 || nObjPts != nImgPts {
		panic("Invalid dimensions for objPts and imgPts")
	}

	objPtsVec := NewGcvPoint3f32Vector(int64(nObjPts))
	imgPtsVec := NewGcvPoint2f32Vector(int64(nObjPts))

	for j := 0; j < nObjPts; j++ {
		objPtsVec.Set(j, NewGcvPoint3f32(mat64.Col(nil, j, objPts)...))
	}

	for j := 0; j < nObjPts; j++ {
		imgPtsVec.Set(j, NewGcvPoint2f32(mat64.Col(nil, j, imgPts)...))
	}

	_imgSize := NewGcvSize2i(dims[0], dims[1])

	camMat = GcvMatToMat64(GcvInitCameraMatrix2D_(
		objPtsVec, imgPtsVec, _imgSize, aspectRatio))
	return camMat
}

func GcvCalibrateCamera(objPts, imgPts, camMat, distCoeffs *mat64.Dense,
	dims [2]int, flags int) (calCamMat, rvec, tvec *mat64.Dense) {

	objDim, nObjPts := objPts.Dims()
	imgDim, nImgPts := imgPts.Dims()

	if objDim != 3 || imgDim != 2 || nObjPts != nImgPts {
		panic("Invalid dimensions for objPts and imgPts")
	}

	objPtsVec := NewGcvPoint3f32Vector(int64(nObjPts))
	imgPtsVec := NewGcvPoint2f32Vector(int64(nObjPts))

	for j := 0; j < nObjPts; j++ {
		objPtsVec.Set(j, NewGcvPoint3f32(mat64.Col(nil, j, objPts)...))
	}

	for j := 0; j < nObjPts; j++ {
		imgPtsVec.Set(j, NewGcvPoint2f32(mat64.Col(nil, j, imgPts)...))
	}

	_camMat := Mat64ToGcvMat(camMat)
	_distCoeffs := Mat64ToGcvMat(distCoeffs)
	_rvec := NewGcvMat()
	_tvec := NewGcvMat()
	_imgSize := NewGcvSize2i(dims[0], dims[1])

	GcvCalibrateCamera_(
		objPtsVec, imgPtsVec,
		_imgSize, _camMat, _distCoeffs,
		_rvec, _tvec, flags)

	calCamMat = GcvMatToMat64(_camMat)
	rvec = GcvMatToMat64(_rvec)
	tvec = GcvMatToMat64(_tvec)

	return calCamMat, rvec, tvec
}

// GcvRodrigues takes a 3D column vector, and apply cv::Rodrigues to it.
func GcvRodrigues(src *mat64.Dense) (dst *mat64.Dense) {
	gcvSrc := Mat64ToGcvMat(src)
	gcvDst := NewGcvMat()
	GcvRodrigues_(gcvSrc, gcvDst)
	dst = GcvMatToMat64(gcvDst)

	return dst
}
