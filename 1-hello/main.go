package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) > 1 {
		fmt.Println("Hello", args[1])
	} else {
		fmt.Println("Hello world")
	}
}
