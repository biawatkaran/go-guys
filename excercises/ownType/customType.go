package main

import "fmt"

type x int

var y x
var z int

func main() {

	s := "hey"

	s = "hello"
	fmt.Println(s)

	y = 23
	fmt.Println(y)
	fmt.Printf("%T\t%+v\n", y, y)

	y = 24
	fmt.Println(y)

	z = int(y)
	fmt.Println(z)
}
