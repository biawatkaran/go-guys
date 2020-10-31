package main

import "fmt"

type contactInfo struct {
	email    string
	postcode string
}

/** to understand with basics
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
**/

// in real world
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	/** To show one way of assinging values to struct
		oneWay := person{"John", "Smith"}
		fmt.Println("One way --> person{\"John\", \"Smith\"} : ", oneWay)
	**/

	preferredWay := person{firstName: "John", lastName: "Smith"}
	fmt.Println("Preferred way --> person{firstName :\"John\", lastName : \"Smith\"} : ", preferredWay)

	var byDefault person
	fmt.Println("Default values : ", byDefault)
	fmt.Printf("%+v", byDefault) //layout the full structure of variable including fields

	byDefault.firstName = "John"
	byDefault.lastName = "Smith"
	fmt.Printf("%+v", byDefault)

	fmt.Println("\nWeird complex nested structs needs comma on every line... duh")
	complexStructDef := person{
		firstName: "complex",
		lastName:  "struct",
		contactInfo: contactInfo{
			email:    "something@gmail.com",
			postcode: "whatever",
		},
	}

	fmt.Println("With % v layout the full structure of variable with plain values excluding field names")
	fmt.Printf("%v", complexStructDef)
	fmt.Println("\nWith % + v layout the full structure of variable including field names")
	fmt.Printf("%+v", complexStructDef)

	fmt.Println("\nCalling in receiver format like structVar.method()")
	complexStructDef.print()

	// ************ Pass By Value ************
	fmt.Println("\n\nPass by Value")
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:    "jim@gmail.com",
			postcode: "address",
		},
	}

	fmt.Println("\tPrinting Normal Value")
	jim.print()

	fmt.Println("\n\tUpdated Directly Value - ofcourse will change")
	jim.firstName = "Jimmy"
	jim.print()

	fmt.Println("\n\tUpdated Value via method - pass by value - won't change")
	jim.updateFirstName("Jim PassByValue")
	jim.print()

	// ************ Here comes pointers ************
	fmt.Println("\n\tUpdated Using Pointers long format - now will change")
	jimPointer := &jim
	jimPointer.updateFirstNameViaPointer("Jim Pointer Long Format")
	jim.print()

	fmt.Println("\n\tUpdated Using Pointers short format - now will change")
	jim.updateFirstNameViaPointer("Jim Pointer Short Format")
	jim.print()
}

func (p person) updateFirstName(newFirstName string) {

	p.firstName = newFirstName
}

func (p *person) updateFirstNameViaPointer(newFirstName string) {

	(*p).firstName = newFirstName
}

func (p person) print() {

	fmt.Printf("\t\t%+v", p)
}
