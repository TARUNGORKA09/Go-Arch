package main

import (
	"fmt"
	"time"
)

func Print(s string) {

	for _, j := range s {
		fmt.Println(string(j))
	}
}

func main() {
	go Print("TARUN")
	time.Sleep(5 * time.Second)
	Print("Gorka")

}
