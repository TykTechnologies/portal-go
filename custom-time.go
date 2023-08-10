package portal

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02 15:04", s)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format("2006-01-02 15:04")), nil
}
