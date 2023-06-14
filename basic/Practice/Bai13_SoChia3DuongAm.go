package main

import "fmt"

func main() {
	var a, b, c int
	var count3, countD, countA int
	fmt.Println("Nhap vao a b c")
	fmt.Scanln(&a, &b, &c)
	if a >= 0 {
		countD++
		if a%3 == 0 {
			count3++
		}
	} else {
		countA++
		if a%3 == 0 {
			count3++
		}
	}
	if b >= 0 {
		countD++
		if b%3 == 0 {
			count3++
		}
	} else {
		countA++
		if b%3 == 0 {
			count3++
		}
	}
	if c >= 0 {
		countD++
		if c%3 == 0 {
			count3++
		}
	} else {
		countA++
		if c%3 == 0 {
			count3++
		}
	}
	fmt.Printf("Co %d so duong\n", countD)
	fmt.Printf("Co %d so am\n", countA)
	fmt.Printf("Co %d so chia het cho 3\n", count3)
}
