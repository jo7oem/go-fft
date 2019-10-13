package fft

import (
	"math"
	"math/cmplx"
)

func DFT(x []complex128) []complex128 {
	lx := len(x)
	res := make([]complex128, lx, lx)

	omega := -1i * 2 * math.Pi
	for i := 0; i < lx; i++ {
		omegaK := omega * complex(float64(i), 0)
		element := x[i]
		for n := 1; n < lx; n++ {
			element += x[n] * cmplx.Exp(omegaK*complex(float64(n), 0))
		}
		res[i] = element
	}
	return res
}

func R2C(x []float64) []complex128 {
	lx := len(x)
	res := make([]complex128, lx, lx)
	for i := 0; i < lx; i++ {
		res[i] = complex(x[i], 0)
	}
	return res
}
func C2R(x []complex128) []float64 {
	lx := len(x)
	res := make([]float64, lx, lx)
	for i := 0; i < lx; i++ {
		res[i] = cmplx.Abs(x[i])
	}
	return res
}
