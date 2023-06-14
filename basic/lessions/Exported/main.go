package main

import (
	"Exported/pers"
	"fmt"
)

func main() {
	p := new(pers.Person)
	// error: p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	//p.firstName = "Eric"
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName()) // Output: Eric
}
