package main

import "fmt"

func main() {
	var tme, tcon int
	fmt.Println("Nhap vao tuoi cua me:")
	fmt.Scan(&tme)
	fmt.Println("Nhap vao tuoi cua con:")
	fmt.Scan(&tcon)
	// tính bao nhiêu năm nữa tuôi con bằng 1/2 tuổi mẹ
	// tme + x = 2*(tcon + x)
	nam := tme - 2*tcon
	fmt.Printf("Sau %v nam nua tuoi con bang 1/2 tuoi me", nam)
}
