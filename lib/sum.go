package lib

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumFromString(stringNumbers []string) (sum int, err error) {
	numbers, err := ConvertStringSliceToIntSlice(stringNumbers)
	if err != nil {
		return 0, err
	}
	return Sum(numbers), nil
}

