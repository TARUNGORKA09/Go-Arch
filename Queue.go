package main

import (
	"fmt"
)

type Queue []int

func (q *Queue) Insert(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pull() int {
	//index = len(*q)-1
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}
func (S *Queue) Print() {
	for _, j := range *S {
		fmt.Println(j)
	}
}

func main() {
	var q Queue
	q.Insert(1)
	q.Insert(2)
	q.Insert(3)
	q.Pull()
	q.Print()

}
