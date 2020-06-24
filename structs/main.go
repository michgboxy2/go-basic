package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimpointer := &jim             //the & gives the memory address of the value this variable is pointing at
	jim.updateName("Jimmy") //update struct
	jim.print()
	// fmt.Printf("%+v", jim) //"%+v" shows key + value whereas "%v" shows only value// {Jim Party {jim@gmail.com 94000}} vs {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}

}

func (pointerToPerson *person) updateName(newFirstName string) { //update *person -: this is a type description, it means we arw working with a pointer to a person.
	(*pointerToPerson).firstName = newFirstName //* give me the value this memory address is pointing at
	//*pointerToPerson -: This is an operator, it means we want to manipulate the value the pointer is referencing.
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)

}
