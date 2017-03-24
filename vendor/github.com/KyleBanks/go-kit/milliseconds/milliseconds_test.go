package milliseconds

import (
	"testing"
	"time"
)

func TestFrom(t *testing.T) {
	res := From(time.Now())
	if res <= 0 {
		t.Error("Unexpected time returned: ", res)
	}
}

func TestNow(t *testing.T) {
	now := Now()
	if now <= 0 {
		t.Error("Unexpected time returned: ", now)
	}
}
