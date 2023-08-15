package main

import (
	"fmt"
)

func main() {
	//var s3 = [...]string{1: "Hello", 0: "World"}
	//var s3 = [...]string{"Hello!", "World"}
	//var a = [...]int{1, 2, 3}
	/*// a là một array

	// b là một con trỏ tới array a
	var b = &a
	// in ra hai phần tử đầu tiên của array a
	fmt.Println(a[0], a[1])
	// truy xuất các phần tử của con trỏ array cũng giống như truy xuất các phần tử của array
	fmt.Println(b[0], b[1])
	// duyệt qua các phần tử trong con trỏ array, giống như duyệt qua array
	for index, value := range b {
		// thay đổi từng phần tử trong b
		b[index] += 1
		fmt.Println(index, value)
	}
	// giá trị của các phần tử trong a bị thay đổi vì b
	for index, value := range a {
		fmt.Println(index, value)
	}*/
	/*for i := range s3 {
		fmt.Printf("a[%d]: %s\n", i, s3[i])
	}*/
	s := "hello world"
	//fmt.Println("len(s): ", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println(len(s))
}
