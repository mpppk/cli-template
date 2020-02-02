package domain

import "github.com/mpppk/cli-template/pkg/util"

type Numbers []int

func NewNumbers(nums []int) Numbers {
	return nums
}

func NewNumbersFromStringSlice(strNumbers []string) (Numbers, error) {
	rawNumbers, err := util.ConvertStringSliceToIntSlice(strNumbers)
	if err != nil {
		return nil, err
	}
	return NewNumbers(rawNumbers), nil
}

func (n Numbers) CalcSum() int {
	return util.Sum(n)
}

func (n Numbers) CalcL1Norm() int {
	return util.L1Norm(n)
}
