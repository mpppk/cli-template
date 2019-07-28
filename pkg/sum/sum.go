// Package sum provides utilities for calculate sum of numbers
package sum

import (
	"github.com/mpppk/cli-template/pkg/util"
	"math"
)

// Sum returns sum of numbers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// L1Norm returns L1 norm of numbers
func L1Norm(numbers []int) (l1norm int) {
	var absNumbers []int
	for _, number := range numbers {
		absNumbers = append(absNumbers, int(math.Abs(float64(number))))
	}
	return Sum(absNumbers)
}

// SumFromString returns sum numbers which be converted from strings
func SumFromString(stringNumbers []string) (sum int, err error) {
	numbers, err := util.ConvertStringSliceToIntSlice(stringNumbers)
	if err != nil {
		return 0, err
	}
	return Sum(numbers), nil
}
