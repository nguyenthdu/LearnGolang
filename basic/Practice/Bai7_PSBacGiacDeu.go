package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("Nhap vao canh a: ")
	fmt.Scan(&a)
	p := 8 * a
	s := (1 + math.Sqrt(2)) * 2 * a * a
	fmt.Println("Chu vi cua hinh bat giac deu la: ", p)
	fmt.Println("Chu vi cua hinh bat giac deu la: ", s)

}
