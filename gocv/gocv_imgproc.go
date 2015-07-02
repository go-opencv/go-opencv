package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
// #cgo linux  pkg-config: opencv
import "C"
import "github.com/gonum/matrix/mat64"

// GcvThreshold takes a 3D column vector, and apply cv::Threshold to it.
func GcvThreshold(src *mat64.Dense) (dst *mat64.Dense, rtn float64) {
	gcvSrc := Mat64ToGcvMat(src)
	gcvDst := NewGcvMat()
	rtn = GcvThreshold_(gcvSrc, gcvDst, 1.0, 2.0, 0)
	dst = GcvMatToMat64(gcvDst)

	return dst, rtn
}
