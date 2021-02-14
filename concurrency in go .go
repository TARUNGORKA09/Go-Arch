package main

import (
	"fmt"
	"time"
)

func main(){

	message := make(chan string)
	go Print("Tarun",message)
	for messages:= range message{
		fmt.Println("Message := ",messages)
	}	
	time.Sleep(1 * time.Second)
}

func Print(t string, message chan string) {
	for i:=1;i<5;i++{
		message <- t
		//time.Sleep(500 * time.Millisecond)
	}
	close(message)
}