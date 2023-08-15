package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	a, b, c float64
}

func semiPerimeter(t Triangle) float64 {
	return (t.a + t.b + t.c) / 2
}

func area(t Triangle) float64 {
	p := semiPerimeter(t)
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))

}

func main() {
	t1 := Triangle{12, 2, 2}
	t2 := Triangle{9, 4, 1}
	fmt.Println("Area of r1 is: ", area(t1))
	fmt.Println("Area of r2 is: ", area(t2))
}
