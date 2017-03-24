// Package today provides utilities for access regarding today's date.
package today

import (
	"time"
)

// BeforeMidnight returns the current date with time set to directly before midnight.
// For example, 2016-06-24 11:59:59.999
func BeforeMidnight() time.Time {
	year, month, day := time.Now().Date()

	return time.Date(year, month, day, 23, 59, 59, 999999999, time.Local)
}
