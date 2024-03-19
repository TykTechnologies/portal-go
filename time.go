// Copyright 2024 Tyk Technologies
// SPDX-License-Identifier: MPL-2.0

package portal

import (
	"strings"
	"time"
)

type Time struct{ time.Time }

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	newTime, err := time.Parse("2006-01-02 15:04", s)
	if err != nil {
		return err
	}

	t.Time = newTime

	return nil
}

func (t *Time) MarshalJSON() []byte {
	return []byte(t.Time.Format("2006-01-02 15:04"))
}
