package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//// Easy syntax but is dependent on the order the structs fields are defined
	//tyler1 := person{"Tyler", "Reece"}
	//fmt.Println(tyler1)
	//
	//// Makes it easier if you want to change order of structs fields
	//tyler2 := person{firstName: "Tyler", lastName: "Reece"}
	//fmt.Println(tyler2)
	//
	//// Assigns "zero" values (falsy values)
	//var tyler3 person
	//fmt.Println(tyler3)
	//fmt.Printf("%+v", tyler3) // Prints with key values
	//
	//tyler3.firstName = "taylor"
	//tyler3.lastName = "rice"
	//
	//fmt.Printf("%+v", tyler3) // Prints with key values

	jim := person{
		firstName: "jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 1,
		}}

	jim.updateName("jimmy")
	jim.print()

	jim.updateNamePointer("jimmy")
	jim.print()
}

// When using Value Types (int, float, string, bool, structs) you have to get the value from the pointer object
func (p *person) updateNamePointer(newFirstName string) {
	(*p).firstName = newFirstName
}

// When using Reference Types (slices, maps, channels, pointers, functions) you can update the value directly
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
