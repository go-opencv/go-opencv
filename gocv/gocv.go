package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
import "C"

func NewGcvPoint3f(x, y, z float32) GcvPoint3f_ {
	return NewGcvPoint3f_(float32(x), float32(y), float32(z))
}

func NewGcvPoint3d(x, y, z float64) GcvPoint3d_ {
	return NewGcvPoint3d_(float64(x), float64(y), float64(z))
}

func NewGcvPoint2f(x, y float32) GcvPoint2f_ {
	return NewGcvPoint2f_(float32(x), float32(y))
}

func NewGcvPoint2d(x, y float64) GcvPoint2d_ {
	return NewGcvPoint2d_(float64(x), float64(y))
}

func NewGcvSize2f(x, y float32) GcvSize2f_ {
	return NewGcvSize2f_(float32(x), float32(y))
}

func NewGcvSize2d(x, y float64) GcvSize2d_ {
	return NewGcvSize2d_(float64(x), float64(y))
}
