package main

import (
	"fmt"
	//"sync"
)

func main() {

	c := make(chan int)
	//var wg sync.WaitGroup

	//wg.Add(10)

	for i := 0; i < 10; i++ {

		go func() {

			c <- 10
		}()
		//wg.Done()
		// close(c) You can't as first run already closed the channel so second run it will close of closed channel :D
	}

	//wg.Wait()

	//close(c) // bit odd as we never know it may close channel and other 10 routines never fill up data at all
	// and following for range return empty, hence used concrete for loop
	/*for i := range c {
		fmt.Println(i)
	}*/

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("About to exit")

}
