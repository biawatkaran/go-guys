package main

import "fmt"

func main() {

	//define and initialize map
	colors := map[string]string{

		"red":   "#ff0000",
		"green": "#ff2306",
	}

	fmt.Println(colors)

	//another way of defining map
	var basicDef map[string]string
	fmt.Println(basicDef)

	//one more way of defining map using function make
	shortMapDef := make(map[string]string)
	fmt.Println(shortMapDef)

	//add values in a map
	shortMapDef["key1"] = "value1"
	shortMapDef["key2"] = "value2"
	fmt.Println(shortMapDef)

	//remove value from map using delete function and pass key
	delete(shortMapDef, "key1")
	fmt.Println(shortMapDef)
}
