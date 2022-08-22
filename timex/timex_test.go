package timex

import "testing"

func TestTimex(t *testing.T) {
	now := Now()
	t.Log(now.ToTime())
	t.Log(now.Ts10())
	t.Log(now.Ts13())
	t.Log(now.Fmt())
}
