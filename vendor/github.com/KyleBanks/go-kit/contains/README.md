# contains
--
    import "github.com/KyleBanks/go-kit/contains/"

Package contains adds a few small helper functions to see if a slice contains a
particular value.

## Usage

#### func  Int

```go
func Int(val int, arr []int) bool
```
Int returns true if the slice of ints contains the value provided.

#### func  Uint

```go
func Uint(val uint, arr []uint) bool
```
Uint returns true if the slice of uints contains the value provided.
