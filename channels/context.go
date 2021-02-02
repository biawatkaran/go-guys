package main

/**
Contexts are mainly used in scenarios where
	you have routines running in the background
	when main thread gets cancelled
	then you like those backgrounds thread to get killed as well to release resource
	that's where context comes handy

**/

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("#routines: ", runtime.NumGoroutine())

	go func() {

		n := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}

	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("#routines: ", runtime.NumGoroutine())

	fmt.Println("About to cancel")
	cancel()
	fmt.Println("already cancelled")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("#routines: ", runtime.NumGoroutine())

}
