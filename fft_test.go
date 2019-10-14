package goFFT

import (
	dspFFT "github.com/mjibson/go-dsp/fft"
	"math"
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
func genWave(lens int) []float64 {
	res := make([]float64, lens, lens)
	flens := float64(lens)
	for i := 0; i < lens; i++ {
		res[i] = math.Sin(1 * 2 * math.Pi * (float64(i) / flens))
	}

	return res

}
func TestDFT(t *testing.T) {
	tmp := genWave(2048)
	wont := dspFFT.FFTReal(tmp)
	res := DFT(R2C(tmp))
	for i, f := range res {
		t.Logf("%d,%f,%f,%f\n", i, tmp[i], f, wont[i])
	}
}
