package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	defer fmt.Println("2")
	defer fmt.Println("1")
	defer fmt.Println("0")
}
