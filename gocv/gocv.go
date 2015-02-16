package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
import "C"
import "github.com/gonum/matrix/mat64"

func NewGcvPoint3f32(x, y, z float64) GcvPoint3f32_ {
	return NewGcvPoint3f32_(float32(x), float32(y), float32(z))
}

func NewGcvPoint3f64(x, y, z float64) GcvPoint3f64_ {
	return NewGcvPoint3f64_(float64(x), float64(y), float64(z))
}

func NewGcvPoint2f32(x, y float64) GcvPoint2f32_ {
	return NewGcvPoint2f32_(float32(x), float32(y))
}

func NewGcvPoint2f64(x, y float64) GcvPoint2f64_ {
	return NewGcvPoint2f64_(float64(x), float64(y))
}

func NewGcvSize2f32(x, y float64) GcvSize2f32_ {
	return NewGcvSize2f32_(float32(x), float32(y))
}

func NewGcvSize2f64(x, y float64) GcvSize2f64_ {
	return NewGcvSize2f64_(float64(x), float64(y))
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
			if fltPtr, ok := mat.GcvAtf64(i, j).(*float64); ok {
				data = append(data, *fltPtr)
			} else {
				panic("Non *float64 passed to MatToMat64")
			}

		}
	}

	return mat64.NewDense(row, col, data)
}
