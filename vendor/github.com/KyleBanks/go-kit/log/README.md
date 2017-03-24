# log
--
    import "github.com/KyleBanks/go-kit/log/"

Package log provides a simple logging service to print to stdout/stderr with
timestamp and log source information.

## Usage

```go
var (
	// Logger to be passed around as a LogWriter instance.
	Logger logger
)
```

#### func  Error

```go
func Error(a ...interface{})
```
Error outputs to stderr.

#### func  Errorf

```go
func Errorf(format string, a ...interface{})
```
Errorf outputs a formatted error to stderr.

#### func  Info

```go
func Info(a ...interface{})
```
Info outputs to stdout.

#### func  Infof

```go
func Infof(format string, a ...interface{})
```
Infof outputs a formatted string to stdout.

#### func  PrintStack

```go
func PrintStack()
```
PrintStack outputs the current go routine's stack trace.
