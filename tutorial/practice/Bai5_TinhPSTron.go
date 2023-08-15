package main

import "fmt"

func main() {
	var r float64
	fmt.Println("Nhap vao ban kinh hinh tron:")
	fmt.Scan(&r)
	p := 2 * 3.14 * r
	s := 3.14 * r * r
	fmt.Printf("Chu vi hinh tron la: %v\n", p)
	fmt.Printf("Dien tich hinh tron la: %.2f\n", s)
}
