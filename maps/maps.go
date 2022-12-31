package main

import (
	"fmt"
)

func main() {

	//maps can be made using the make function
	lastWorldSeriesWinners := make(map[string]int)
	lastWorldSeriesWinners["Astros"] = 2022
	lastWorldSeriesWinners["Braves"] = 2021
	lastWorldSeriesWinners["Dodgers"] = 2020

	//or maps can be made and instantiated like this
	worldSeriesTitles := map[string]int{
		"Yankees":      27,
		"Cardinals":    11,
		"Red Sox":      9,
		"Athletics":    9,
		"Giants":       8,
		"Dodgers":      7,
		"Pirates":      5,
		"Reds":         5,
		"Braves":       4,
		"Tigers":       4,
		"White Sox":    3,
		"Twins":        3,
		"Orioles":      3,
		"Cubs":         3,
		"Phillies":     2,
		"Guardians":    2,
		"Mets":         2,
		"Astros":       2,
		"Royals":       2,
		"Blue Jays":    2,
		"Marlins":      2,
		"Diamondbacks": 1,
		"Angels":       1,
		"Nationals":    1,
		"Padres":       0,
		"Rangers":      0,
		"Rays":         0,
		"Brewers":      0,
		"Rockies":      0,
		"Mariners":     0,
	}

	printMap(lastWorldSeriesWinners)

	//updating, adding, and deleting from a map
	worldSeriesTitles["Cardinals"] = 12
	worldSeriesTitles["Expansion Team"] = 0
	fmt.Println("After WS win and expansion team added")
	printMap(worldSeriesTitles)

	delete(worldSeriesTitles, "Expansion Team")
	fmt.Println("Expansion team removed")
	printMap(worldSeriesTitles)

	//maps are pass by reference, so cheap to pass around
	addExpansionTeam(worldSeriesTitles, "Ducks")
	printMap(worldSeriesTitles)

	//Additional Notes on Maps:
	//Best practice is to specify map size, especially for larger maps
	//Maps are not thread safe, so not safe for concurrency

}

func printMap(incomingMap map[string]int) {
	//ranging over a map will give unique order
	for key, val := range incomingMap {
		fmt.Printf("Key is: %v Value is: %v\n", key, val)
	}
}

func addExpansionTeam(incomingMap map[string]int, teamName string) {
	incomingMap[teamName] = 0
}
