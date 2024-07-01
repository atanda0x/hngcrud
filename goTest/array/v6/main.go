package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(numberOfSum ...[]int) []int {
	var sum []int

	for _, numbers := range numberOfSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, Sum(tail))
		}
	}

	return sum
}
