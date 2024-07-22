package main

import "fmt"

const (
	WIDTH  = 1920 // columns of 2D array
	HEIGHT = 1080 // rows of 2D array
)

type pixel int                  // aliasing int as pixel
var screen [WIDTH][HEIGHT]pixel // global 2D array

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0 // initializing value to 2D array
		}
	}
	//print screen
	fmt.Print(screen)
}
