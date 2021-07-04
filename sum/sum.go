package sum

func sum(numbers []int) (ret int) {

	for _, number := range numbers {
		ret += number
	}
	return
}

func sumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}
	return
}

func sumAllTails(numbersToSumTails ...[]int) (sums []int) {

	for _, numbers := range numbersToSumTails {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			numbers = numbers[1:]
			sums = append(sums, sum(numbers))
		}

	}
	return
}
