package main

import (
	"fmt"
	"strconv" //string converter
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Sam", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	person2 := Person{"Bill", "Nye", "Singapore", "m", 75}
	// fmt.Println(person2.firstName)
	// person1.age++
	// fmt.Println(person1)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	person1.getMarried("Yogster")

	person2.getMarried("Jenner")

	fmt.Println(person1.greet())

}
