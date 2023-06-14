package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x1, x2, delta float64
	fmt.Println("Nhap vao a b c ")
	fmt.Scanln(&a, &b, &c)
	delta = b*b - 4*a*c
	if delta > 0 {
		x1 = (-b - math.Sqrt(delta)) / 2 * a
		x2 = (-b + math.Sqrt(delta)) / 2 * a
		fmt.Printf("PT co 2 nghiem x1 = %f, x2 = %f", x1, x2)
	} else if delta == 0 {
		x1 = -b / (2 * b)
		fmt.Printf("PT co 1 nghiem duy nhat x1=", x1)
	} else {
		fmt.Println("Phuong trinh vo nghiem")
	}
}
