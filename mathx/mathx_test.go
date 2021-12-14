package mathx

import "testing"

func TestSum(t *testing.T) {

	t.Log(Sum(1, 2, 3, 4.4, 5.5, 6.6))
}

func TestMax(t *testing.T) {
	t.Log(Max(1, 2, 3, 4.4, 5.5, 6.6))
	t.Log(Max(-1, -2, -3, -4.4, -5.5, -6.6))
}

func TestMin(t *testing.T) {
	t.Log(Min(1, 2, 3, 4.4, 5.5, 6.6))
	t.Log(Min(-1, -2, -3, -4.4, -5.5, -6.6))
}
