package main

import (
	"fmt"
	"math"
)

func main() {

	var arr = []float64{4.0, 5.0, 1.0, 9.0, 2.0}
	var max_no float64 = arr[0]
	var min_no float64 = arr[0]
	for i, _ := range arr {
		max_no = math.Max(max_no, arr[i])
		min_no = math.Min(min_no, arr[i])
	}
	fmt.Println(max_no)
	fmt.Println(min_no)

}
