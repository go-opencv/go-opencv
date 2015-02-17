package gocv

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gonum/matrix/mat64"
)

func TestNewGcvPoint3f32(t *testing.T) {
	pt := NewGcvPoint3f32(3, 1)
	spew.Dump(pt)
}

func TestNewGcvPoint2f32(t *testing.T) {
	pt := NewGcvPoint2f32(3, 1)
	spew.Dump(pt)
}

func TestNewGcvSize2f64(t *testing.T) {
	size := NewGcvSize2f64(3, 1)
	spew.Dump(size)
}

func TestMat(t *testing.T) {
	mat := NewGcvMat()
	mat2 := NewGcvMat(mat)
	spew.Dump(mat2)
}

func TestMat64ToGcvMat(t *testing.T) {
	mat := mat64.NewDense(3, 2, []float64{
		0, 1,
		1.23, 4,
		-12.3, -4,
	})
	spew.Dump(Mat64ToGcvMat(mat))
}
