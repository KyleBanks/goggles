# timer
--
    import "github.com/KyleBanks/go-kit/timer/"

Package timer provides the ability to time abritrary events, like the duration
of a method call.

## Usage

#### func  New

```go
func New() func() time.Duration
```
New returns a timer that can be used to measure the duration between the call to
New() and the call to the function returned by new.

For example:

    t := New()
    SomeOtherStuff()
    duration := t()

#### func  NewLogger

```go
func NewLogger(msg string) func()
```
NewLogger returns a timer that immediately logs 'BEGAN <msg>' and again logs
'ENDED <msg>' when the returned function is executed, along with the time
between the two.

Intended to be used as so:

    l := NewLogger("loading")
    defer l()

Prints:

    BEGAN loading
    ENDED loading 1234 ms
