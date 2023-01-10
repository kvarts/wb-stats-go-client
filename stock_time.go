package openapi

import (
	"strings"
	"time"
)

type StockTime struct {
	time.Time
}

func (c *StockTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}

	c.Time, err = time.Parse("2006-01-02T15:04:05", s)
	return
}
