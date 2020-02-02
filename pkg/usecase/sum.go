package usecase

import (
	"github.com/mpppk/cli-template/pkg/domain"
	"golang.org/x/xerrors"
)

// CalcL1SumFromStringSlice is use case to calculate sum from string slice
func CalcSumFromStringSlice(strNumbers []string) (int, error) {
	numbers, err := domain.NewNumbersFromStringSlice(strNumbers)
	if err != nil {
		return 0, xerrors.Errorf("failed to create numbers from string slice: %w", err)
	}
	return numbers.CalcSum(), nil
}

// CalcL1NormFromStringSlice is use case to calculate L1 norm from string slice
func CalcL1NormFromStringSlice(strNumbers []string) (int, error) {
	numbers, err := domain.NewNumbersFromStringSlice(strNumbers)
	if err != nil {
		return 0, xerrors.Errorf("failed to create numbers from string slice: %w", err)
	}
	return numbers.CalcL1Norm(), nil
}
