package main

import "fmt"

type Student struct {
	name string
	age  int
}

func f1(a, b, c int) int { // taking three parameters and returning their sum
	return a + b + c
}

func f2(a, b int) (int, int, int) { // taking two parameters and returning their sum, difference and product
	n1 := a + b
	n2 := a - b
	n3 := a * b
	return n1, n2, n3
}
func inputInfo(name string, age int) Student {
	return Student{name, age}

}

type (
	IZ int
	FZ float64
)

func main() {
	fmt.Print(f1(f2(20, 10))) // function call within a function call
	fmt.Println()
	//sử dụng kiểu dữ liệu  đã định nghĩa
	var a IZ
	var b FZ
	a = 5
	b = 5.5
	fmt.Println(a)
	fmt.Println(b)
}
