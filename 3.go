package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	//scanner.Split(bufio.ReadWords)

	for scanner.Scan() {
		lines := scanner.Text()
		fmt.Println(lines)
		fmt.Fprintf(conn, "I Heard That you say : %s \n\n", lines)
	}
}
