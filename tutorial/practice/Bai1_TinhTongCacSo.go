package main

import "fmt"

func main() {
	var n, m, a, b, c, t int
	fmt.Print("Nhap vao so nguyen co 3 chu so: ")
	//Phương thức lấy giá trị n nhập vào từ bàn phím
	fmt.Scan(&n)  //123
	a = n / 100   // lấy chữ số hàng trăm: 1
	m = n % 100   //lấy đươc 2 chữ số hàng đơn vị: 23
	b = m / 10    //lấy được 2
	c = m % 10    //lấy được 3
	t = a + b + c //tổng 3 chữ số
	fmt.Println("Tong cac chu so cua so nguyen la: ", t)
}
