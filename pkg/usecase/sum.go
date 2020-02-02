package usecase

import (
	"github.com/mpppk/cli-template/pkg/domain"
	"golang.org/x/xerrors"
)

func CalcSumFromStringSlice(strNumbers []string) (int, error) {
	numbers, err := domain.NewNumbersFromStringSlice(strNumbers)
	if err != nil {
		return 0, xerrors.Errorf("failed to create numbers from string slice: %w", err)
	}
	return numbers.CalcSum(), nil
}

func CalcL1NormFromStringSlice(strNumbers []string) (int, error) {
	numbers, err := domain.NewNumbersFromStringSlice(strNumbers)
	if err != nil {
		return 0, xerrors.Errorf("failed to create numbers from string slice: %w", err)
	}
	return numbers.CalcL1Norm(), nil
}
