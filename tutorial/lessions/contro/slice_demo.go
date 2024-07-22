package main

import "fmt"

func main() {
	//s := make([]int, 5)
	//for i := range s {
	//	s[i] = i
	//}
	//for _, v := range s {
	//	println(v)
	//}
	//s1 := []int{1, 2, 3, 4, 5}
	//s2 := s1[1:3]
	//for _, v := range s2 {
	//	println(v)
	//}
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, 5)
	//copy(dest, src)
	n := copy(dest, src)
	fmt.Println(n)
	fmt.Println(dest)
}
