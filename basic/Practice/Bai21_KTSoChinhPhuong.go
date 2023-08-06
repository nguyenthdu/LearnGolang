package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Nhap vao so can kiem tra: ")
	fmt.Scanln(&n)
	m := int(math.Sqrt(float64(n)))
	if m*m == n {
		fmt.Printf("%d la so chinh phuong", n)
	} else {
		fmt.Printf("%d khong phai la so chinh phuong", n)
	}
}
