package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestTimex(t *testing.T) {
	now := Now()
	t.Log(now.ToTime())
	t.Log(now.Ts10())
	t.Log(now.Ts13())
	t.Log(now.Fmt())
}

func TestForTs(t *testing.T) {
	now := Now()
	ts := now.Ts13()
	suffix := ts % 10000
	other := time.Unix(ts/1000, suffix)

	fmt.Println(other)
}
