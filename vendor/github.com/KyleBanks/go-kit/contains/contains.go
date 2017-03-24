// Package contains adds a few small helper functions to see if a
// slice contains a particular value.
package contains

// Int returns true if the slice of ints contains the value provided.
func Int(val int, arr []int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

// Uint returns true if the slice of uints contains the value provided.
func Uint(val uint, arr []uint) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}
