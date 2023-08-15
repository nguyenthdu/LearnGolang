package main

import "fmt"

func main() {
	var a, h float64
	fmt.Println("Nhap vao canh va chieu cao hinh thoi:")
	fmt.Scanln(&a, &h)
	p := 4 * a
	s := a * h
	fmt.Printf("Chu vi hinh thoi la: %f\nChu vi hinh thoi la: %f", p, s)

}
