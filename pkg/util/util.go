package util

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/xerrors"
)

func ConvertStringSliceToIntSlice(stringSlice []string) (intSlice []int, err error) {
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, xerrors.Errorf("failed to convert string slice to int slice: %w", err)
		}
		intSlice = append(intSlice, num)
	}
	return
}

func EPrintlnIFErrExist(err error, msg string) (bool, error) {
	if err != nil {
		if _, err := fmt.Fprintln(os.Stderr, msg); err != nil {
			return true, err
		}
		return true, nil
	}
	return false, nil
}
