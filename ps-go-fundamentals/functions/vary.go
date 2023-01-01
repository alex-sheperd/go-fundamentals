package main

import (
	"fmt"
)

func main() {
	averageTemp := calculateAverage(77, 88, 67, 92, 84, 86)
	fmt.Println(averageTemp)
}

func calculateAverage(temps ...int) int {
	totalTemps := 0
	for _, i := range temps {
		totalTemps += i
	}
	return totalTemps / len(temps)
}
