package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Nhap vao 3 so a b c:")
	fmt.Scanln(&a, &b, &c)
	var max int
	if a > b && a > c {
		max = a
	} else if b > a && b > c {
		max = b
	} else {
		max = c
	}
	fmt.Printf("so lon nhat trong 3 so %d %d %d la: %d", a, b, c, max)

}
