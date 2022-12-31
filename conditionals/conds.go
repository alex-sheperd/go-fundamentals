package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	//if we declare our variables inside the if statement, then they are scoped
	//to that conditonal and are garbage collected after the condition executes

	if myAge, yourAge := 34, 27; myAge > yourAge {
		fmt.Println("I am older")
	} else if myAge == yourAge {
		fmt.Println("We are the same age")
	} else {
		fmt.Println("You are older")
	}

	//a declared value in the switch will also be garbage colleted
	//the break is implicit to each scenario
	//fallthrough can be used to have one case spill into the next
	switch "The Fellowship of the Ring" {
	case "The Fellowship of the Ring":
		fmt.Println("First")
	case "The Two Towers":
		fmt.Println("Second")
	case "The Return of the King":
		fmt.Println("Third")
	default:
		fmt.Println("Prequel")
	}

	//idiomatic switch statements have multiple cases to avoid using fallthroughs

	switch tempNum := rng(); tempNum {
	case 1, 2, 3, 5, 7, 11:
		fmt.Println("The number is prime:", tempNum)
	case 0, 4, 6, 8, 9, 10, 12:
		fmt.Println("The number is not prime", tempNum)
	}

	//if statements can also be used to check for error handling
	_, err := os.Open("./test/txt")

	if err != nil {
		fmt.Println("There is the error code:", err)
	}
}

func rng() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(13)
}
