package main

import (
	"fmt"
)

func main() {

	type user struct {
		firstName string
		lastName  string
		age       int
	}

	//various creation methods
	//var alex user
	//alex := new(user)

	//create with initial values
	alex := user{
		firstName: "Alex",
		lastName:  "Sheperd",
		age:       34,
	}

	//accessors
	fmt.Println("My name is", alex.firstName, alex.lastName, "and I am", alex.age, "years old")

	//modifiers
	alex.age = 35
	fmt.Println("Next year I will be", alex.age)
}
