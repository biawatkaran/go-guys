package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 1. a type satisfies interface, when it implements all the functions contains in interface definition
type logWriter struct{}

func readSiteData() {

	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	//long way of doing it
	fmt.Println("********** Long form to display **********")
	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))

	//short way of doing same
	fmt.Println("********** Short form to display **********")
	resp, err = http.Get("http://www.google.com")
	io.Copy(os.Stdout, resp.Body)

	//short way of doing same
	fmt.Println("********** Custom form to display **********")
	lw := logWriter{}
	resp, err = http.Get("http://www.google.com")
	io.Copy(lw, resp.Body)
}

//2. writer interface defines Write function and our custom type implements it, so logWriter satisfies Writer interface
func (logWriter) Write(bs []byte) (n int, e error) {

	fmt.Println(string(bs))
	fmt.Printf("Wrote #bytes to stdout %v", len(bs))

	return len(bs), nil
}
