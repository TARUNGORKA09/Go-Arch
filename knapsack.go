package main

import "fmt"

func main() {
	var value = []int{60, 100, 120}
	var weight = []int{10, 20, 30}
	w := 50
	res := knapsack(value, weight, len(value), w)
	fmt.Println(res)
}

func knapsack(val []int, wt []int, n int, w int) int {

	if n == 0 || w == 0 {
		return 0
	}
	if wt[n-1] <= w {
		return val[n-1] + knapsack(val, wt, n-1, w-wt[n-1])
	}
	return knapsack(val, wt, w, n-1)

}
