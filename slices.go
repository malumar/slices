package mslices

import "fmt"

// Unique remove duplicated values from array
// @return new table with unique values
func Unique[T comparable](ar []T) []T {
	if ar == nil {
		return nil
	}
	ret := make([]T, 0)

	for i := range ar {
		found := false
		for j := range ret {
			if i == j {
				continue
			}

			if ar[i] == ret[j] {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, ar[i])
		}
	}

	return ret
}

// Filter filters items using the keep function
// @return return new array with filtered values
func Filter[T any](ar []T, keep func(val T) bool) []T {
	if ar == nil {
		return nil
	}
	ret := make([]T, 0, len(ar))
	for _, x := range ar {
		if keep(x) {
			ret = append(ret, x)
		}
	}
	return ret
}

// SliceOfFlag array support for variables in program flags
// np.
// var list sliceOfFlag[string]
// flag.Var(&list, "item", "additional items")
type SliceOfFlag[T any] []T

// String is an implementation of the flag.Value interface
func (i *SliceOfFlag[T]) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set is an implementation of the flag.Value interface
func (i *SliceOfFlag[T]) Set(value T) error {
	*i = append(*i, value)
	return nil
}
