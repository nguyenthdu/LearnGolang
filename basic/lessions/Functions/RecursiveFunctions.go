package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i) // function call
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

// Function to calculate the nth fibonacci number
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2) // recursive call
	}
	return
}
