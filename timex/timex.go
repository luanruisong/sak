package timex

import (
	"time"
)

type Time time.Time

func (t Time) ToTime() time.Time {
	return time.Time(t)
}
func (t Time) Ts10() int64 {
	return t.ToTime().Unix()
}
func (t Time) Ts13() int64 {
	return t.ToTime().UnixNano() / int64(time.Millisecond)
}
func (t Time) Fmt() string {
	return t.ToTime().Format("2006-01-02 15:04:05")
}

func Now() Time {
	return Time(time.Now())
}

func ForTs10(ts int64) Time {
	return Time(time.Unix(ts, 0))
}

func ForTs13(ts int64) Time {
	return ForTs10(ts / 1000)
}
