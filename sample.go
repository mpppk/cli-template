package main

import (
	"errors"
	"fmt"
)

func IncrementPositiveNumber(v int) (int, error) {
	if v == 0 {
		err := errors.New("v is zero")
		return 0, err
	}

	if v < 0 {
		return 0, fmt.Errorf("v is negative number: %d", v)
	}
	return v + 1, nil
}
