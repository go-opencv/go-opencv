package gocv

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewGcvPoint3f32(t *testing.T) {
	pt := NewGcvPoint3f32(3, 1, 2)
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
	mat := NewMat()
	mat2 := NewMat(mat)
	spew.Dump(mat2)
}
