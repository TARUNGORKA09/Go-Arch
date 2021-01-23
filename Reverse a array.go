package main

import (
	"fmt"
)

func main() {

	var arr = []float64{4.0, 5.0, 1.0, 3.0, 9.0, 2.0}
	Reverse(arr)

}

func Reverse(arr []float64) {

	i := 0
	j := len(arr) - 1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
		if i == j {
			break
		}
	}
	for _, j := range arr {
		fmt.Println(j)
	}

}
