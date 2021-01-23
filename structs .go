package main

import "fmt"

type Student struct {
	name  string
	class int
}

func main() {

	var t Student
	t.name = "TARUN"
	t.class = 10

	fmt.Println(&t)

}
