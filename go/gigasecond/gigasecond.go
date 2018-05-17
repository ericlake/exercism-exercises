package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond takes a start time and adds 10 ^ 9 seconds to it
func AddGigasecond(t time.Time) time.Time {
	duration := math.Pow(10, 9)

	return t.Add(time.Duration(int(duration)) * time.Second)
}
