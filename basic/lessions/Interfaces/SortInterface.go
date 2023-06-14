package main

import (
	"LearnGolang/basic/lessions/Interfaces/mysort"
	"fmt"
)

// sorting of slice of integers
/*
Hàm ints1 sắp xếp một slice của các số nguyên.
Đầu tiên, slice data chứa các số nguyên được khai báo.
Tiếp theo, slice data được chuyển đổi thành kiểu mysort.
IntSlice để sử dụng trong package mysort.
Hàm mysort.Sort được gọi để sắp xếp slice data theo thứ tự tăng dần.
Cuối cùng, kiểm tra xem slice đã được sắp xếp chưa và in ra kết quả.
*/
func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := mysort.IntSlice(data) //conversion to type IntSlice
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

// sorting of slice of strings1
/*
Hàm strings1 sắp xếp một slice của các chuỗi.
Tương tự như hàm ints1, slice data chứa các chuỗi được khai báo và được chuyển đổi thành kiểu mysort.StringSlice.
Hàm mysort.Sort được gọi để sắp xếp slice data theo thứ tự từ điển.
Cuối cùng, kiểm tra xem slice đã được sắp xếp chưa và in ra kết quả.
*/
func strings() {
	data := []string{"Monday", "Friday", "Tuesday", "Wednesday", "Sunday", "Thursday", "", "Saturday"}
	a := mysort.StringSlice(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

// a type which describes a day of the week
type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

// sorting of custom type day
func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	mysort.Sort(&a)
	if !mysort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()
}
