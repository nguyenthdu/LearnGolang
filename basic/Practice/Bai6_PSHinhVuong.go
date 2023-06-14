package main

import (
	"fmt"
)

const PI = 3.14

func main() {
	var a float64
	fmt.Println("Nhap vao canh a: ")
	fmt.Scan(&a)
	p := 4 * a
	s := PI * a * a
	fmt.Println("Chu vi hinh vuong la: ", p)
	fmt.Println("Dien tich hinh vuong la: ", s)
}
