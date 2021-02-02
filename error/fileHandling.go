package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("somerandomfile.txt")

	if err != nil {

		fmt.Println("Err: ", err)
		log.Println("Err: ", err)
		return
	}

	defer f.Close()

	r := strings.NewReader("Wassup")

	io.Copy(f, r)
}
