package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex //remove this mu references and you have race condition check using `go run -race main.go` command

	for i :=0; i<gs; i++ {

		go func(){
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("final counter:", counter)
}
