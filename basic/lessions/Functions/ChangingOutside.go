package main

import (
	"fmt"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b // side-effect(changing n)
}

func main() {
	n := 0
	reply := &n // một biến con trỏ mới replyvà lưu địa chỉ của n
	//n va reply dùng chung 1 địa chỉ nên chỉ cần 1 trong 2 thay đổi thì cái còn lại cũng thay đổi
	fmt.Println("Before Multiplication:", n)
	fmt.Println("Before Multiplication:", *reply)

	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("Multiply:", n)      // Multiply: 50
}
