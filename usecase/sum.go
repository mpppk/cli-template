package usecase

import (
	"fmt"

	"github.com/mpppk/cli-template/domain"
)

// CalcSum is use case to calculate sum
func CalcSum(strNumbers []int) int {
	return domain.NewNumbers(strNumbers).CalcSum()
}

// CalcSumFromStringSlice is use case to calculate sum from string slice
func CalcSumFromStringSlice(strNumbers []string) (int, error) {
	numbers, err := domain.NewNumbersFromStringSlice(strNumbers)
	if err != nil {
		return 0, fmt.Errorf("failed to create numbers from string slice: %w", err)
	}
	return numbers.CalcSum(), nil
}

// CalcL1Norm is use case to calculate L1 norm
func CalcL1Norm(strNumbers []int) int {
	return domain.NewNumbers(strNumbers).CalcL1Norm()
}

// CalcL1NormFromStringSlice is use case to calculate L1 norm from string slice
func CalcL1NormFromStringSlice(strNumbers []string) (int, error) {
	numbers, err := domain.NewNumbersFromStringSlice(strNumbers)
	if err != nil {
		return 0, fmt.Errorf("failed to create numbers from string slice: %w", err)
	}
	return numbers.CalcL1Norm(), nil
}
