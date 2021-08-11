package rdb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

var (
	ErrInvalidDBValueForTime = errors.New("invalid db value for time")
)

type Time struct {
	time time.Time
}

func FromTime(t time.Time) Time {
	return Time{time: t}
}

func (t *Time) ToTime() time.Time {
	return t.time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return t.time.MarshalJSON()
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var s time.Time
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t.time = s
	return nil
}

func (t Time) String() string {
	return t.time.String()
}

// Scan implements the Scanner interface.
func (t *Time) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}
	var ok bool
	t.time, ok = value.(time.Time)
	if !ok {
		return ErrInvalidDBValueForTime
	}
	return nil
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return t.time.Format(timeLayout), nil
}

//func (t Time) Equal(ct Time) bool {
//	return t.time.Equal(t.time)
//}
