package main

import "fmt"

func sum(a []int) int { // function that sums integers
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4} // declare an array
	fmt.Println(sum(arr[:]))        // passing slice to the function
}
