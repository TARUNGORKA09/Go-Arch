package main

import "fmt"

func main() {

	var set = []int{1, 5, 5, 11}
	var sum int
	for _, j := range set {
		sum += j
	}
	if sum%2 == 1 {
		fmt.Println("false")
	} else {
		res := EqualSumPart(set, sum/2, len(set))
		fmt.Println(res)
	}
}

func EqualSumPart(set []int, sum int, n int) bool {
	if sum == 0 || n == 0 {
		return false
	}
	if set[n-1] > sum {
		return EqualSumPart(set, sum, n-1)
	}
	return EqualSumPart(set, sum-set[n-1], n-1) || EqualSumPart(set, sum, n-1)
}
