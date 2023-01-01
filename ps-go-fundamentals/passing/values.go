package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Alex"

	//go passes by value by default
	// we can use pointers to pass by reference
	//& gives us the memory address of the variable
	fmt.Println("Memory address of the *name* variable is", &name)

	//the * before the variable type means it's a pointer variable and returns the value stored in the variable it's being pointed to
	var pointer *string = &name

	//here we use the * to dereference the pointer to get the value
	fmt.Println("The address of the pointer variable is", pointer, "which has the value of", *pointer)

	//Here we can change the value
	fmt.Println("Hello, my name is", name)
	updateName(name)

	//this does not update the name because we passed name by value, therefore we changed a copy of the variable name
	//and not to the actual name variable
	fmt.Println("I passed my name by value, my name is now", name)

	//pass the name by reference to update.
	pointerUpdateName(&name)
	fmt.Println("I passed my name by reference, my name is now", name)

	//we can access environment variables via the OS package
	envUser := os.Getenv("USER")
	fmt.Println("Environment User:", envUser)

}

func updateName(name string) string {
	name = "Alexander"
	fmt.Println("Name updated to", name)
	return name
}

func pointerUpdateName(name *string) string {
	*name = "Alexander"
	return *name
}
