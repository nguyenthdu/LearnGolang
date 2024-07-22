package main

import "fmt"

func main() {
	//var x float64 = 1.7e+308
	//fmt.Printf("Type: %T, value: %.2f", x, x)
	//LABEL:
	//	for i := 0; i < 10; i++ {
	//		for j := 0; j < 10; j++ {
	//			if j == 2 {
	//				continue LABEL
	//			}
	//			fmt.Printf("i: %d, j: %d\n", i, j)
	//		}
	//
	//	}
	i := 0
HERE: // adding a label for a location
	fmt.Printf("%d", i)
	i++
	if i == 5 {
		return
	}
	goto HERE // goto HERE
}
