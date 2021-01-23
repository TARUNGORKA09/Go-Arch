package main

import (
	"fmt"
)

func main() {

	var arr = []int{2, 6, 4, 5, 9}

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-2; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			i++
		}
	}
	for _, j := range arr {
		fmt.Println(j)
	}

}
