package main

import "fmt"

func main() {
	var a, b float64
	fmt.Println("Nhap vao chieu dai và chieu rong")
	fmt.Scanf("%f %f", &a, &b)
	p := 2 * (a + b)
	s := a * b
	fmt.Println("Chieu dai cua hinh chu nhat la: ", p)
	fmt.Println("Chieu rong cua hinh vuong la: ", s)
}
