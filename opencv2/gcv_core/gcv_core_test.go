package gcv_core

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewGcvPoint3f(t *testing.T) {
	pt := NewGcvPoint3f(3, 1, 2)
	spew.Dump(pt)
}

func TestNewGcvPoint2f(t *testing.T) {
	pt := NewGcvPoint2f(3, 1)
	spew.Dump(pt)
}

func TestNewGcvSize2d(t *testing.T) {
	size := NewGcvSize2d(3, 1)
	spew.Dump(size)
}
