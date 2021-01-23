package main

import "fmt"

type student struct {
	roll     int
	subjects Subjects
	hobbies Hobbies
}

type Subjects struct {
	Eng   bool
	Hindi bool
}
type Hobbies struct {
	Cricket bool
}






func main() {

	rahul := []*student{
		{
			roll:     1,
			subjects: Subjects{Eng: true, Hindi: true},
		},
		{
			roll:     2,
			subjects: Subjects{Eng: true, Hindi: false},
		},
	}

	fmt.Print(rahul[0])

	fmt.Print(rahul[1])

}
