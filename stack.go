package main

import "fmt"

type stack []int

func (s *stack) Push(t int) {
	*s = append(*s, t)
}
func (S *stack) IsEmpty() bool {
	return len(*S) == 0
}
func (S *stack) Pull() (int, bool) {
	if S.IsEmpty() {
		return 0, false
	} else {
		index := len(*S) - 1
		element := (*S)[index]
		*S = (*S)[:index]
		return element, true

	}
}
func (S *stack) Print() {
	for _, j := range *S {
		fmt.Println(j)
	}
}

func (S *stack) Top() int {
	index := len(*S) - 1
	return (*S)[index]
}

func main() {
	var S stack
	S.Push(1)
	S.Push(2)
	S.Push(3)
	S.Push(4)
	S.Print()
}
