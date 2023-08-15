package main

import "fmt"

func getData() (int, string, bool) {
	return 10, "Hello", true
}

func main() {

	_, _, message := getData()
	fmt.Println("Message:", message)
}
