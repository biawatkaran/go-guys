package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	//receive - range method keeps on reading from c, the final value again waiting so has to close c
	// range blocks until you close the channel
	for v := range c {

		fmt.Println(v)
	}

	fmt.Println("close testing")
}
