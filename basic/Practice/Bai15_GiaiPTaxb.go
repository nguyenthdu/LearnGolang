package main

import "fmt"

func main() {
	var a, b, x float64
	fmt.Println("Nhập vào 2 số  a và b: ")
	fmt.Scanln(&a, &b)
	if a <= 0 {
		fmt.Println("Chuong trinh loi a phai lon hon 0")

	} else {
		x = -b / a
		fmt.Println("Gia tri cua x la: ", x)
	}
}
