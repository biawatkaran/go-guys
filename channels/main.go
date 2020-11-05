package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{

		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // make a channel of type string here, could be of any type

	for _, link := range links {

		fmt.Println("checking...", link)
		go checklink(link, c) //without channel it won't hold the main routine, so w/o waiting go routine main will finish hence channels
	}

	/**
	//blocking waits for data to come on channel as soon it gets, then prints it
	//1.
	fmt.Println(<- c)

	//2. call as many links - simple hack'ish
	for i:=0; i<len(links); i++ {
		fmt.Println(<-c)
	}

	func checklink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println("link might be down - ", link)
		c <- "Might be down I think!"
		os.Exit(1)
	}

	fmt.Println("link is working ", link)
	c <- "Yep its up!"
	}

	//3. call repeatedly say infinite
	for { //short `for i := 0; i < len(links); i++` to run infinite
		go checklink(<-c, c) // note the first value, channel internally infers its of type string so let it go. its kind of not
	                        // that friendly as user would have to figure out what coming into <-c and scroll through code :(
	}
	**/

	//4. call repeatedly as in spamming someone :P
	for l := range c { //more understandable way from developer perspective aka watch c to spit value and once come out, use loop
		//time.Sleep(5 * time.Second) // Dont put here, as it blocking aka sleeping for 5sec even if c has got data back from others
		//go checklink(l, c)

		//we call functional literal instead which is like anonymous methods
		go func(link string) {
			time.Sleep(3 * time.Second)
			checklink(link, c)
		}(l)
	}
}

func checklink(link string, c chan string) {

	//time.Sleep(5 * time.Second) // Dont put here, now every child routine waits 5 seconds before starting
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("link might be down - ", link)
		c <- link //note the value return here is link, check the commented part in 2. #41 #46
		return
	}

	fmt.Println("link is working ", link)
	c <- link
}
