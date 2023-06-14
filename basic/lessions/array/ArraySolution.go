package main

import "fmt"

func main() {
	//var fib [10]int64
	fmt.Println(fibs())
}

func fibs() (fib [10]int64) {
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < len(fib); i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}
