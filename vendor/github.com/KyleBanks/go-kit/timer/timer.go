// Package timer provides the ability to time abritrary events, like the duration
// of a method call.
package timer

import (
	"log"
	"time"
)

// New returns a timer that can be used to measure the duration
// between the call to New() and the call to the function returned by new.
//
// For example:
// 		t := New()
//		SomeOtherStuff()
// 		duration := t()
func New() func() time.Duration {
	t := time.Now()
	return func() time.Duration {
		return time.Now().Sub(t)
	}
}

// NewLogger returns a timer that immediately logs 'BEGAN <msg>'
// and again logs 'ENDED <msg>' when the returned function is executed,
// along with the time between the two.
//
// Intended to be used as so:
//		l := NewLogger("loading")
//		defer l()
//
// Prints:
//		BEGAN loading
//		ENDED loading 1234 ms
func NewLogger(msg string) func() {
	t := New()
	log.Printf("BEGAN %v", msg)

	return func() {
		log.Printf("ENDED %v  %v", msg, t())
	}
}
