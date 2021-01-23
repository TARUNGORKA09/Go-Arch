package main

import (
	"fmt"
	"io"
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
		io.WriteString(conn, "Hello World")
		fmt.Println("Tarun Gorka")
		conn.Close()
	}

}
