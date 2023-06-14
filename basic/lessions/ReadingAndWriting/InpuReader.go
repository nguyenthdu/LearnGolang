package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input1 string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input1: ")

	input1, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input1 was: %s\n", input1)
	}
}
