package main

import (
	"fmt"
	"math"
	"reflect"
)

var (
	name, course  string
	module, clip  int
	inferredName  = "This is a string"
	inferredInt   = 4
	inferredFloat = 4.0
)

func main() {
	fmt.Println("Variables")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("The variable inferredName can be inferred as a type", reflect.TypeOf(inferredName))
	fmt.Println("The variable inferredInt can be inferred as a type", reflect.TypeOf(inferredInt))

	//you cannot add different types
	//you have to convert the types
	//here we can convert a float to a int using the round functionality
	sum := inferredInt + int(math.Round(inferredFloat))
	fmt.Println("Our inferred float added to our int returns", sum)

	//variables declared inside a method must be called or you will get a runtime error.
	// variables declared at the global scope do not need to be used
	userActive := false
	fmt.Println("User active:", userActive)

}
