package main

import (
	"fmt"
)

func main() {

	var arr = []int{2, 6, 5, 5, 4, 9, 4, 5, 9, 4, 6, 5}

	var count int = 0
	for _, j := range arr {
		if j == 5 {
			count++
		}
	}
	fmt.Println(count)

}
