package main

import "fmt"

func main() {

	var set = []int{3, 34, 4, 12, 5, 2}
	var sum int = 9
	result := subset_sum_problem(set, len(set), sum)
	fmt.Println(result)

}

func subset_sum_problem(set []int, n int, sum int) bool {
	if sum == 0 || n == 0 {
		return false
	}
	if set[n-1] > sum {
		return subset_sum_problem(set, n-1, sum)
	}
	return subset_sum_problem(set, n-1, sum-set[n-1]) || subset_sum_problem(set, n-1, sum)
}
