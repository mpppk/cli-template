package sum

import (
	"github.com/mpppk/cli-template/pkg/util"
	"math"
)

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func L1Norm(numbers []int) (l1norm int) {
	var absNumbers []int
	for _, number := range numbers {
		absNumbers = append(absNumbers, int(math.Abs(float64(number))))
	}
	return Sum(absNumbers)
}

func SumFromString(stringNumbers []string) (sum int, err error) {
	numbers, err := util.ConvertStringSliceToIntSlice(stringNumbers)
	if err != nil {
		return 0, err
	}
	return Sum(numbers), nil
}
