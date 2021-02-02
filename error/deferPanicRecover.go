package main

import "fmt"

func main() {

	f()
	fmt.Println("Returned normally from f()")
}

func f() {

	defer func() {

		if r := recover(); r != nil { // recover method has to be inside defer
			fmt.Println("Recovered in f() with value ", r) // recover handles the panic and stops this method
			// further execution #24 won't run
			// make it resumes as normal in original caller here main()
		}
	}()

	fmt.Println("Calling g()")
	g(0)

	fmt.Println("Returned normally from g()")
}

func g(i int) {

	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in g()", i) // runs in LIFO order, runs before panic goes to the caller
	fmt.Println("Printing in g() ", i)
	g(i + 1)
}
