# today
--
    import "github.com/KyleBanks/go-kit/today/"

Package today provides utilities for access regarding today's date.

## Usage

#### func  BeforeMidnight

```go
func BeforeMidnight() time.Time
```
BeforeMidnight returns the current date with time set to directly before
midnight. For example, 2016-06-24 11:59:59.999
