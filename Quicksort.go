package main

import (
	"fmt"
)

func main() {

	var arr = []int{38, 27, 43, 3, 9, 82, 10}
	m := 0
	r := len(arr) - 1
	quick_sort(arr, r, m)
	Print(arr)
}
func quick_sort(arr []int, r int, m int) {
	l := m
	h := r
	for l < h {
		if arr[l] > arr[h] {
			arr[l], arr[h] = arr[h], arr[l]
			l++
		} else {
			h--
		}
	}
	arr[l], arr[r-1] = arr[r-1], arr[l]
}

func Print(arr []int) {
	for _, j := range arr {
		fmt.Println(j)
	}
}
