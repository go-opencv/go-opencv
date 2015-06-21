package gocv

// #cgo CXXFLAGS: -std=c++11
// #cgo darwin pkg-config: opencv
// #cgo linux  pkg-config: opencv
import "C"
import "github.com/gonum/matrix/mat64"

func NewGcvPoint3f32(pts ...float64) GcvPoint3f32_ {
	// This make sure we have default values
	safePts := getSafePts(pts, 3)
	return NewGcvPoint3f32_(float32(safePts[0]), float32(safePts[1]), float32(safePts[2]))
}

func NewGcvPoint3f64(pts ...float64) GcvPoint3f64_ {
	safePts := getSafePts(pts, 3)
	return NewGcvPoint3f64_(safePts[0], safePts[1], safePts[2])
}

func NewGcvPoint2f32(pts ...float64) GcvPoint2f32_ {
	safePts := getSafePts(pts, 2)
	return NewGcvPoint2f32_(float32(safePts[0]), float32(safePts[1]))
}

func NewGcvPoint2f64(pts ...float64) GcvPoint2f64_ {
	safePts := getSafePts(pts, 2)
	return NewGcvPoint2f64_(safePts[0], safePts[1])
}

func NewGcvSize2f32(pts ...float64) GcvSize2f32_ {
	safePts := getSafePts(pts, 2)
	return NewGcvSize2f32_(float32(safePts[0]), float32(safePts[1]))
}

func NewGcvSize2f64(pts ...float64) GcvSize2f64_ {
	safePts := getSafePts(pts, 2)
	return NewGcvSize2f64_(safePts[0], safePts[1])
}

// Convert Mat, which defined by SWIG, to *mat64.Dense.
// The reason is the latter is much easier to handle
// in Go.
// GcvMat is assumed to be 2-dimensional matrix.
func GcvMatToMat64(mat GcvMat) *mat64.Dense {
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

// Convert *mat64.Dense to Mat
func Mat64ToGcvMat(mat *mat64.Dense) GcvMat {
	row, col := mat.Dims()

	rawData := NewGcvFloat64Vector(int64(row * col))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			rawData.Set(i*col+j, mat.At(i, j))
		}
	}

	return Mat64ToGcvMat_(row, col, rawData)
}

func getSafePts(pts []float64, size int) []float64 {
	// This make sure we have default values
	safePts := make([]float64, size, size)
	copy(safePts, pts)

	return safePts
}
