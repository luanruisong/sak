package mathx

import (
	"testing"
)

func TestSumFloat(t *testing.T) {
	t.Log(SumFloat(1.1, 2.2, 3.3))
}

func TestMaxFloat(t *testing.T) {
	t.Log(MaxFloat(1.1, 2.2, 3.3))
	t.Log(MaxFloat(-1.1, -2.2, -3.3))
}

func TestMinFloat(t *testing.T) {
	t.Log(MinFloat(1.1, 2.2, 3.3))
	t.Log(MinFloat(-1.1, -2.2, -3.3))
}
