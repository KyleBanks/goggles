// Package convert provides generalized type conversion utilities.
package convert

import (
	"errors"
	"fmt"
	"strconv"
)

// StringSliceToIntSlice accepts a slice of strings and returns a slice of parsed ints.
//
// If any of the strings cannot be parsed to an integer, an error will be returned.
func StringSliceToIntSlice(strs []string) ([]int, error) {
	ints := make([]int, len(strs), len(strs))

	for i, str := range strs {
		anInt, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			return nil, errors.New("Failed to parse UserId[" + str + "] due to: " + err.Error())
		}

		ints[i] = int(anInt)
	}

	return ints, nil
}

// IntSliceToStringSlice converts a slice of ints to a slice of strings.
func IntSliceToStringSlice(ints []int) []string {
	strings := make([]string, len(ints), len(ints))

	for i, v := range ints {
		strings[i] = strconv.Itoa(v)
	}

	return strings
}

// UintSliceToStringSlice converts a slice of uints to a slice of strings.
func UintSliceToStringSlice(ints []uint) []string {
	strings := make([]string, len(ints), len(ints))

	for i, v := range ints {
		strings[i] = fmt.Sprintf("%v", v)
	}

	return strings
}

// SliceToStringSlice converts a slice into a slice of their string representations.
func SliceToStringSlice(s []interface{}) []string {
	strings := make([]string, len(s), len(s))

	for i, v := range s {
		strings[i] = fmt.Sprintf("%v", v)
	}

	return strings
}
