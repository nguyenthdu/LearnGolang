package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("Nhap vao canh a: ")
	fmt.Scan(&a)
	p := 6 * a
	s := 3 * math.Sqrt(3) * a * a / 2
	fmt.Printf("Chu vi hinh luc giac la: %v\n", p)
	fmt.Printf("Dien tich hinh luc giac la: %v\n", s)
}
