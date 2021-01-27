package main

import "fmt"

func main() {

	var arr = []int{5, 1, 4, 2, 8}
	Bubble_sort(arr)
	Print(arr)
}

func Bubble_sort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

func Print(arr []int) {
	for _, j := range arr {
		fmt.Println(j)
	}
}
