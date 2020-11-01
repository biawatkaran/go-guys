package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	args := os.Args
	readFile(args[1])
}

func readFile(fileName string) {

	fptr, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fptr)
}
