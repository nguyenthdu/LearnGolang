package main

import (
	"fmt"
)

// const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6
type Gender int

const (
	UNKNOWN Gender = iota
	FAMALE
	MALE
)

func main() {
	fmt.Println(MALE)
}
