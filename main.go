package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func main() {
	// john := person{"John", "Steinbeck", 66}

	// john := person{
	// 	firstName: "John",
	// 	lastName:  "Steinbeck",
	// 	age:       66,
	// }

	var john person
	john.firstName = "John"
	john.lastName = "Steinbeck"
	john.age = 66
	john.contactInfo = contactInfo{
		email:   "john@gmail.com",
		zipCode: "90210",
	}

	// fmt.Println(john.contact.email)
	// fmt.Printf("%+v\n", john)
	// john.print()

	// john.updateName("Jonathan")
	// john.print() //=> "John"

	// johnPointer := &john
	// johnPointer.updateName("Jonathan") // Updates the memory address for "John"
	// john.print()                       //=> "Jonathan"

	// Shortcut to deal with pointers. Golang allows us to call a function that
	// requires a pointer by either the pointer or instance.
	john.updateName("Jonathan")
	john.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
