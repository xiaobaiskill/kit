package rdb

import (
	"database/sql/driver"
	"strings"
)

type Strings []string

func (s Strings) String() string {
	return strings.Join(s, ",")
}

func (s *Strings) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	v, ok := value.([]byte)
	if !ok {
		return ErrInvalidDBValueForStrings
	}
	*s = strings.Split(string(v), ",")
	return nil
}

func (s Strings) Value() (driver.Value, error) {
	return strings.Join(s, ","), nil
}
