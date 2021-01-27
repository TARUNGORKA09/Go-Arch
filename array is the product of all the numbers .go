package main

import "fmt"

func main() {

	var arr = []int{1, 2, 3, 4, 5}
	var mul []int

	i := 0
	var prod_left int
	var prod_right int
	for i < len(arr) {

		for j := i; j > 0; j-- {
			prod_left = arr[j] * arr[j-1]
		}

		for j := i; j < len(arr)-2; j++ {
			prod_right = arr[j] * arr[j+1]
		}
		mul = append(mul, prod_left+prod_right)
		i++
	}
	Print(mul)
}

func Print(mul []int) {
	for _, j := range mul {
		fmt.Println(j)
	}

}
