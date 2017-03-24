// Package job provides the ability to execute tasks on a timed interval.
package job

import (
	"time"
)

// Job is a container for a repeating execution of a function.
//
// A Job executes it's function continuously with a predefined delay between
// executions. The Job only stops when the `Stop` function is called.
type Job struct {
	fn             func()
	delay          time.Duration
	runImmediately bool
	firstRun       bool

	stop    chan bool
	started bool
}

// Stop halts the execution of the Job's function.
func (j *Job) Stop() {
	j.stop <- true
}

// start begins the execution of the Job's function.
//
// Note: It is unsafe to call `start` more than once!
func (j *Job) start() {
	go func() {
		for {
			// Check if runImmediately is set on the first run
			if j.firstRun && j.runImmediately {
				j.fn()
			}
			j.firstRun = false

			// Sleep for the predetermined time.
			time.Sleep(j.delay)

			select {
			// Check for the 'stop' signal.
			case <-j.stop:
				return

			// Execute the function.
			default:
				j.fn()
			}
		}
	}()
}

// Register schedules a function for execution, to be invoked repeated with a delay of
// the value of i.
//
// If the runImmediately parameter is true, the function will execute immediately. Otherwise,
// it will be invoked first after the duration of i.
func Register(f func(), delay time.Duration, runImmediately bool) *Job {
	j := Job{
		fn:             f,
		delay:          delay,
		runImmediately: runImmediately,
		firstRun:       true,

		stop: make(chan bool, 1),
	}

	j.start()
	return &j
}
