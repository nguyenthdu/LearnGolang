package main

import "fmt"

func main() {
	var f, c float64
	fmt.Println("Nhap vao nhiet do F:")
	fmt.Scan(&f)
	c = (f - 32) * 5 / 9
	fmt.Printf("%v do F bang  %.2f do C", f, c)
}
