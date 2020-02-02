package domain

import "github.com/mpppk/cli-template/pkg/util"

// Numbers represents numbers
type Numbers []int

// NewNumbers is constructor for Numbers
func NewNumbers(nums []int) Numbers {
	return nums
}

// NewNumbersFromStringSlice create new Numbers with numbers from string slice
func NewNumbersFromStringSlice(strNumbers []string) (Numbers, error) {
	rawNumbers, err := util.ConvertStringSliceToIntSlice(strNumbers)
	if err != nil {
		return nil, err
	}
	return NewNumbers(rawNumbers), nil
}

// CalcSum calc sum of numbers
func (n Numbers) CalcSum() int {
	return util.Sum(n)
}

// CalcL1Norm calc L1 norm of numbers
func (n Numbers) CalcL1Norm() int {
	return util.L1Norm(n)
}
