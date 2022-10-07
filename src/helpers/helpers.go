package helpers

import (
	"math/rand"
	"time"
)

type Interval struct {
	Start time.Time
	End   time.Time
}

func RandInterval(year int) Interval {
	result := Interval{}

	min := time.Date(year, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(year, 12, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	secEnd := sec + (rand.Int63n(86400) + 1)

	result.Start = time.Unix(sec, 0)
	result.End = time.Unix(secEnd, 0)

	return result
}
