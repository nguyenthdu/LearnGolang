package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "3"
	var nn int
	var err error
	nn, err = strconv.Atoi(orig)
	if err != nil {
		fmt.Print("Error", orig)
		return

	}
	fmt.Print("Success", nn)

}
