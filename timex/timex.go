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
