package mathx

import "testing"

func TestSumInt(t *testing.T) {
	t.Log(SumInt(1, 2, 3))
}

func TestMaxInt(t *testing.T) {
	t.Log(MaxInt(1, 2, 3))
	t.Log(MaxInt(-1, -2, -3))
}

func TestMinInt(t *testing.T) {
	t.Log(MinInt(1, 2, 3))
	t.Log(MinInt(-1, -2, -3))
}
