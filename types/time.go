package types

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	return err
}
