package main

import "fmt"

func main() {

	var t = []int{1, 2, 3, 4, 5, 6}

	for _, j := range t {
		if j%2 == 0 {
			fmt.Print(j)
			fmt.Println("Even")
		} else {
			fmt.Print(j)
			fmt.Println("Odd")
		}
	}
	var First string
	fmt.Scanln(&First)
	var Last string
	fmt.Scanln(&Last)
	fmt.Println("Hello :"+First, Last)

}
