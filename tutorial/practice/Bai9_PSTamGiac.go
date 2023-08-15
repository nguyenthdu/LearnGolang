package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Println("Nhap vao tạo độ x1,y1")
	fmt.Scanln(&x1, &y1)
	fmt.Println("Nhap vao tạo độ x2,y2")
	fmt.Scanln(&x2, &y2)
	fmt.Println("Nhap vao tạo độ x3,y3")
	fmt.Scanln(&x3, &y3)
	a := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	b := math.Sqrt(math.Pow(x2-x3, 2) + math.Pow(y2-y3, 2))
	c := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))
	cv := a + b + c
	p := cv / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Printf("chu vi tam giac la: %.2f\n", cv)
	fmt.Printf("Dien tich tam giac la: %.2f", s)

}
