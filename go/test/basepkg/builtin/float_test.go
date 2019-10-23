package basePkgT

import "testing"

func TestFloatM(t *testing.T) {
	i := 90
	fi := float64(i) * 1.10
	t.Log(fi)
	t.Log(float64(i)*1.10 == float64(99))
}
