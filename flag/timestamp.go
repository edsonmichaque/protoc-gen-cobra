package flag

import (
	"time"
	"unsafe"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TimestampValue struct {
	ptr unsafe.Pointer
}

func NewTimestampValue(value **timestamp.Timestamp) *TimestampValue {
	return &TimestampValue{unsafe.Pointer(value)}
}

func (v *TimestampValue) Set(s string) (err error) {
	for _, fmt := range []string{
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02T15:04:05.999999999Z07",
		"2006-01-02T15:04:05.999999999",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z07",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02T15",
		"2006-01-02",
	} {
		var t time.Time
		t, err = time.Parse(fmt, s)
		if err != nil {
			continue
		}
		*(**timestamp.Timestamp)(v.ptr) = timestamppb.New(t)
		break
	}
	return
}

func (v *TimestampValue) Type() string { return "Timestamp" }

func (v *TimestampValue) String() string { return "<nil>" }