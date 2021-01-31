package main

import (
	//"math"
	"fmt"
)

func Rodcut(length int,price []float64) float64 {

	if length == 0 {
		return 0
	}
	var maxval float64 = 0
	var max float64
	for i:=1;i<length;i++ {
		max = price[i-1]+Rodcut(length-1,price)
	}
	if max>maxval {
		maxval = max
	}
	return maxval
}
func main(){
	var length int = 8
	price := []float64{1,5,8,9,10,17,17,20}
	m := Rodcut(length,price)
	fmt.Println(m)

}