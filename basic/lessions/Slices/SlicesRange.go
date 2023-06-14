package main

import (
	"fmt"
	"strings"
)

func main() {
	slice1 := make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4
	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)

	}
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
	for ix, season := range seasons { // printing seasons
		fmt.Printf("Season %d is: %s \n", ix, season)
	}
	for ix := range seasons {
		seasons[ix] = strings.ToUpper(seasons[ix]) // modifying the seasons
	}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s \n", ix, season) // printing modified seasons
	}
	//==================================
	value := 0            // used to set values to 2D array
	screen := [2][2]int{} // will contain 4 values

	// asssigning values in 2D-array
	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = value
			value = value + 1
		}
	}

	// printing 2D-array
	for row := range screen {
		for column := range screen[0] {
			fmt.Print(screen[row][column], " ") // separates columns
		}
		fmt.Print("\n") //separates rows
	}

}
