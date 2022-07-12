package main

func Sum(n []int) int {
	sum := 0
	for _, number := range n {
		sum += number
	}
	return sum
}

func SumAll(n ...[]int) (sums []int) {
	for _, numbers := range n {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(n ...[]int) (sums []int) {
	for _, numbers := range n {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
