package main

func Sum(n []int) int {
	sum := 0
	for _, number := range n {
		sum += number
	}
	return sum
}
