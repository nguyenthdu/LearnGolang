package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Println("Nhap vao x:")
	fmt.Scanln(&x)
	if x >= 0 {
		y = math.Pow(5, x) + 1
	} else {
		y = (math.Log(-x) + 1) / 2 * x
	}
	fmt.Println("y=", y)
}
