package today

import (
	"testing"
	"time"
)

func TestBeforeMidnight(t *testing.T) {
	year, month, day := time.Now().Date()

	v := BeforeMidnight()

	// Test the date
	if year != v.Year() || month != v.Month() || day != v.Day() {
		t.Fatalf("Date does not match: {Expected: %v-%v-%v, Actual: %v}", year, month, day, v)
	}

	// And the time
	if v.Hour() != 23 || v.Minute() != 59 || v.Second() != 59 || v.Nanosecond() != 999999999 {
		t.Fatalf("Invalid time returned: %v", v)
	}
}
