# convert
--
    import "github.com/KyleBanks/go-kit/convert/"

Package convert provides generalized type conversion utilities.

## Usage

#### func  IntSliceToStringSlice

```go
func IntSliceToStringSlice(ints []int) []string
```
IntSliceToStringSlice converts a slice of ints to a slice of strings.

#### func  SliceToStringSlice

```go
func SliceToStringSlice(s []interface{}) []string
```
SliceToStringSlice converts a slice into a slice of their string
representations.

#### func  StringSliceToIntSlice

```go
func StringSliceToIntSlice(strs []string) ([]int, error)
```
StringSliceToIntSlice accepts a slice of strings and returns a slice of parsed
ints.

If any of the strings cannot be parsed to an integer, an error will be returned.

#### func  UintSliceToStringSlice

```go
func UintSliceToStringSlice(ints []uint) []string
```
UintSliceToStringSlice converts a slice of uints to a slice of strings.
