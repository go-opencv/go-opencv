package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
import "C"
import "github.com/gonum/matrix/mat64"

func NewGcvPoint3f(x, y, z float64) GcvPoint3f_ {
	return NewGcvPoint3f_(float32(x), float32(y), float32(z))
}

func NewGcvPoint3d(x, y, z float64) GcvPoint3d_ {
	return NewGcvPoint3d_(float64(x), float64(y), float64(z))
}

func NewGcvPoint2f(x, y float64) GcvPoint2f_ {
	return NewGcvPoint2f_(float32(x), float32(y))
}

func NewGcvPoint2d(x, y float64) GcvPoint2d_ {
	return NewGcvPoint2d_(float64(x), float64(y))
}

func NewGcvSize2f(x, y float64) GcvSize2f_ {
	return NewGcvSize2f_(float32(x), float32(y))
}

func NewGcvSize2d(x, y float64) GcvSize2d_ {
	return NewGcvSize2d_(float64(x), float64(y))
}

// Convert Mat, which defined by SWIG, to mat64.Dense.
// The reason is the latter is much easier to handle
// in Go.
// GcvMat is assumed to be 2-dimensional matrix.
func MatToMat64(mat Mat) *mat64.Dense {
	col := mat.GetCols()
	row := mat.GetRows()

	data := []float64{}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if fltPtr, ok := mat.GcvAtd(i, j).(*float64); ok {
				data = append(data, *fltPtr)
			} else {
				panic("Non *float64 passed to MatToMat64")
			}

		}
	}

	return mat64.NewDense(row, col, data)
}
