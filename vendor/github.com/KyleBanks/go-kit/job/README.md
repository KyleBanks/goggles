# job
--
    import "github.com/KyleBanks/go-kit/job/"

Package job provides the ability to execute tasks on a timed interval.

## Usage

#### type Job

```go
type Job struct {
}
```

Job is a container for a repeating execution of a function.

A Job executes it's function continuously with a predefined delay between
executions. The Job only stops when the `Stop` function is called.

#### func  Register

```go
func Register(f func(), delay time.Duration, runImmediately bool) *Job
```
Register schedules a function for execution, to be invoked repeated with a delay
of the value of i.

If the runImmediately parameter is true, the function will execute immediately.
Otherwise, it will be invoked first after the duration of i.

#### func (*Job) Stop

```go
func (j *Job) Stop()
```
Stop halts the execution of the Job's function.
