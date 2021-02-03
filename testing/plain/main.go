package main

import "fmt"

func main() {

	fmt.Println("2 +3 =", sum(2, 3))

}

func sum(n ...int) int {

	sum := 0

	for _, v := range n {

		sum += v
	}

	return sum
}
