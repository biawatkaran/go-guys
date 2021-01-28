package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//recieve
	recieve(even, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	//q <- 0 //with this channel gets value 0 and OK idiom says true @43
	close(q) //without this, you get error as q channel needs closing. with this closed, you OK false no value passed @43
}

func recieve(e, o, q <-chan int) {

	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel: ", v)
		case v := <-o:
			fmt.Println("from the odd channel: ", v)
		case v, ok := <-q:
			fmt.Println("from the quit channel: ", v, ok)
			return
		}
	}

}
