package sum

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) (sums []int) {
	for _, slice := range numbers {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(numbers ...[]int) (sums []int) {
	for _, slice := range numbers {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return
}
