package main

import (
	"fmt"
)

func main() {

	//make takes types, here we are making a slice of type string, with a length of 5, and a capacity of 10
	//length is how many entries it has
	//capacity gives the max or the protected size
	//capacity also gives the expected size of the referencing array
	courses := make([]string, 5, 10)

	declaratedSlice := []string{"Item One", "Item Two", "Item Three", "Item Four"}
	myIntSlice := make([]int, 1, 4)

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))

	for _, i := range declaratedSlice {
		fmt.Println(i)
	}

	//when delcaring a slice of a slice, the first number is inclusive, the second is exclusive
	sliceOfSlice := declaratedSlice[1:3]
	for _, i := range sliceOfSlice {
		fmt.Println(i)
	}

	//appending to a slice
	//if the associated array runs out of capacity, under the hood, Go will create a new array with twice the length of the existing
	//and copy all elements to it.
	fmt.Printf("Length of int slice is %d and capacity is %d\n", len(myIntSlice), cap(myIntSlice))

	for i := 1; i < 17; i++ {
		myIntSlice = append(myIntSlice, i)
		fmt.Printf("Length of int slice is %d and capacity is %d\n", len(myIntSlice), cap(myIntSlice))
	}

	anotherSlice := []int{20, 30, 40}
	sliceToAppend := []int{1, 2, 3}

	anotherSlice = append(anotherSlice, sliceToAppend...)
	for _, i := range anotherSlice {
		fmt.Println(i)
	}
}
