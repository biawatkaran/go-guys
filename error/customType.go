package main

import (
	"fmt"
	"log"
)

type customError struct {
	lat  string
	long string
	err  error
}

func (n customError) Error() string {

	return fmt.Sprintf("error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {

	_, err := sqrt(-12)

	if err != nil {
		log.Println(err)
		fmt.Println("Assertion ", err.(customError).err)
	}

}

func sqrt(f float64) (float64, error) {

	if f < 0 {
		ce := fmt.Errorf("oops sqrt of negative number %v", f)
		return 0, customError{"50.5435", "45.46454", ce}
	}

	return 5, nil
}
