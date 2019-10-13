package fft

import (
	"testing"
)

func TestR2C(t *testing.T) {
	r := []float64{4, 3.3, 2.2, -1.1, 0}
	want := []complex128{4, 3.3, 2.2, -1.1, 0}
	res := R2C(r)
	lr := len(r)
	for i := 0; i < lr; i++ {
		if res[i] != want[i] {
			t.Errorf("want = %f, got = %f\n", want[i], res[i])
		}
	}
}
