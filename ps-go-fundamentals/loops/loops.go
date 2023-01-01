package main

import (
	"fmt"
)

func main() {

	// commenting out timer for now
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("BOOM IN SPACE!")
			break
		}
		fmt.Println(timer)
		// time.Sleep(1 * time.Second)
	}

	kings := []string{
		"Joffrey Baratheon",
		"Balon Greyjoy",
		"Rob Stark",
		"Renley Baratheon",
		"Stannis Baratheon"}

	for _, i := range kings {
		fmt.Println(i)
	}

	//breaks with labels can be used to break from different portions of the loop

}
