package main

import (
	"container/list"
)

func Hashcode(key int,size int) int {
	return key % size
}
func main(){
	var HashArr = list.New()
	var Key int = 2
	var Size int = 20
	HCode := Hashcode(Key,Size)

	//for i,j:=range HashArr{
	//	HashArr[i] = list.New()
	//}
	//if HashArr[HCode].value == nil {
	//	HashArr[HCode].value = key
	//}
}