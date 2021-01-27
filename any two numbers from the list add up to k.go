package main

import "fmt"

func main() {

	var arr = []int{10, 3, 7, 15}

	var k int = 17
	sumoftwo(arr, k)
	//fmt.Print(res)
}
func sumoftwo(arr []int, k int) {

	i := len(arr) - 1
	j := 0
	if i == 0 {
		fmt.Println("false")
	} else {
		for j < i {
			if k == arr[i]+arr[j] {
				fmt.Println("true")
			}
			j++
		}
	}
	if i > 0 {
		arr = arr[:i]
		sumoftwo(arr, k)
	}
}
