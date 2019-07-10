package sliceprime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected bool
	}{
		"Empty slices are not prime":   {input: []int{}, expected: false},
		"0 is not prime":               {input: []int{0}, expected: false},
		"1 is prime":                   {input: []int{1}, expected: true},
		"2 is prime":                   {input: []int{1, 1}, expected: true},
		"3 is prime":                   {input: []int{2, 1}, expected: true},
		"4 is not prime":               {input: []int{2, 2}, expected: false},
		"11 is prime":                  {input: []int{2, 5, 4}, expected: true},
		"21 is not prime":              {input: []int{2, 2, 3, 2, 4, 5, 3}, expected: false},
		"46 is not prime":              {input: []int{10, 10, 10, 10, 6}, expected: false},
		"32452867 is prime":            {input: []int{32452867}, expected: true},
		"32452869 is not prime":        {input: []int{32452869}, expected: false},
		"Over MaxInt32 aka 2147483647": {input: []int{715827883, 715827883, 715827883}, expected: false},
	}

	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			actual := IsPrime(tc.input)

			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime([]int{700000001, 700000001, 700000001})
	}
}
