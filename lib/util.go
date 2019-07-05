package lib

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

func ConvertStringSliceToIntSlice(stringSlice []string) (intSlice []int, err error) {
	for _, s := range stringSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert string slice to int slice")
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
