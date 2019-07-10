package sliceprime

import (
	"math"
)

// IsPrime sums all numbers in the provided slice and then determines whether or not that
// sum is a prime number. For cases where errors could be returned, false will be
// returned in their stead until product says otherwise.
func IsPrime(in []int) bool {
	// empty slices
	if len(in) == 0 {
		return false
	}

	num := 0
	for _, n := range in {
		if num+n > math.MaxInt32 {
			return false
		}
		num = num + n
	}

	// zero
	if num == 0 {
		return false
	}

	// handle 1, 2, 3
	if num <= 3 {
		return true
	}

	// eliminate even numbers
	if num%2 == 0 {
		return false
	}

	// eliminate odd numbers
	srf := math.Floor(math.Sqrt(float64(num)))
	sr := int(srf)
	for i := 3; i <= sr; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}

	// it's prime if it's made it this far
	return true
}
