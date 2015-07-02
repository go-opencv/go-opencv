package gocv

import (
	"testing"

	"github.com/gonum/matrix/mat64"
	"github.com/stretchr/testify/assert"
)

func TestGcvThreshold(t *testing.T) {
	rvec := mat64.NewDense(3, 1, []float64{
		3,
		-0.3,
		0.2,
	})
	rmat, _ := GcvThreshold(rvec)

	assert.InDeltaSlice(t, []float64{0.59922526, 0.57799222, -0.55394411},
		rmat.Row(nil, 0), 1e-5)
	assert.InDeltaSlice(t, []float64{0.20413818, 0.558743, 0.80382452},
		rmat.Row(nil, 1), 1e-5)
	assert.InDeltaSlice(t, []float64{0.77411672, -0.5947531, 0.21682264},
		rmat.Row(nil, 2), 1e-5)
}
