package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	var x string = "Hello World/n Tarun gorka/n"

	Scanner := bufio.NewScanner(strings.NewReader(x))

	for Scanner.Scan() {

		line := Scanner.Text()
		fmt.Print(line)

	}

}
