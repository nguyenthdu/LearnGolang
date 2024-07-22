package main

import "fmt"

type Set struct {
	m map[int]bool
}

func NewSet() *Set { //khởi tạo và trả về 1 set rỗng
	return &Set{make(map[int]bool)}

}
func (s *Set) Add(value int) {
	s.m[value] = true
}
func (s *Set) Remove(value int) {
	delete(s.m, value)
}
func (s *Set) Contains(value int) bool {
	return s.m[value]
}
func (s *Set) Elements() []int {
	ele := []int{}
	for key, _ := range s.m {
		ele = append(ele, key)
	}
	return ele
}

func main() {
	// Khởi tạo Set
	set := NewSet()

	// Thêm phần tử vào Set
	set.Add(1)
	set.Add(2)
	set.Add(2)

	// Kiểm tra sự tồn tại của phần tử
	fmt.Println("Set chứa '1':", set.Contains(1)) // true
	fmt.Println("Set chứa '2':", set.Contains(2)) // false

	// Xóa phần tử khỏi Set
	//set.Remove(3)

	// Lấy tất cả các phần tử trong Set
	fmt.Println("Các phần tử trong Set:", set.Elements())
}
