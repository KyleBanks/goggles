package milliseconds

import (
	"time"
)

// From converts a Time to milliseconds since epoch.
func From(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// Now returns the current Time to milliseconds since epoch.
func Now() int64 {
	return From(time.Now())
}
