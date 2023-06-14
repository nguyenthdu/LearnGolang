package main

import "fmt"
import "log"

func main() {
	log.SetFlags(log.Llongfile)
	var where = log.Print

	where()
	// some code
	a := 2 * 5
	where()

	// some more code
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
}
