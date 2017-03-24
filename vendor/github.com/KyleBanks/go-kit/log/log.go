// Package log provides a simple logging service to print to stdout/stderr with timestamp
// and log source information.
package log

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

type logger struct {
}

var (
	// Logger to be passed around as a LogWriter instance.
	Logger logger

	gopath = os.Getenv("GOPATH") + "/src/"
)

// Println is a wrapper for the Info log method.
//
// This is used for passing the Logger around as a LogWriter interface.
func (l logger) Print(a ...interface{}) {
	Info(a)
}

// Info outputs to stdout.
func Info(a ...interface{}) {
	fmt.Println(fileAndLineNumber(), dateAndTimeStamp(), a)
}

// Infof outputs a formatted string to stdout.
func Infof(format string, a ...interface{}) {
	// Note: Cannot call Info() directly because that would ruin the file/line number of the caller
	fmt.Println(fileAndLineNumber(), dateAndTimeStamp(), fmt.Sprintf(format, a...))
}

// Error outputs to stderr.
func Error(a ...interface{}) {
	fmt.Fprintln(os.Stderr, fileAndLineNumber(), dateAndTimeStamp(), a)
}

// Errorf outputs a formatted error to stderr.
func Errorf(format string, a ...interface{}) {
	// Note: Cannot call Error() directly because that would ruin the file/line number of the caller
	fmt.Fprintln(os.Stderr, fileAndLineNumber(), dateAndTimeStamp(), fmt.Sprintf(format, a...))
}

// PrintStack outputs the current go routine's stack trace.
func PrintStack() {
	debug.PrintStack()
}

// fileAndLineNumber returns the file and line number of the code that called `log`.
func fileAndLineNumber() string {
	// Use 2 for the Caller because 0 is this function,
	// 1 is the log.* method that called it,
	// and 2 is what came before.
	_, fn, line, _ := runtime.Caller(2)

	return fmt.Sprintf("%v:%v", strings.Replace(fn, gopath, "", 1), line)
}

// dateAndTimeStamp returns a formatted date/time stamp for the current time.
func dateAndTimeStamp() string {
	return time.Now().Format(time.StampMilli)
}
