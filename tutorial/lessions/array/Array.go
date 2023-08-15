package main

import "fmt"

func main() {
	// khai bao mang
	var arr1 [5]int
	//var arrLazy = [...]int{5, 6, 7, 8, 22}
	//var arrKeyValue = [5]string{3: "Chris", 4: "Ron"} xac dinh ro vi tri truyen tham so vao
	//Một mảng trong Go là một kiểu giá trị (không phải là một con trỏ tới phần tử đầu tiên như trong C/C++). Do đó, nó có thể được tạo bằng new():
	//var arr2 = new([5]int)
	/*TODO:
	Nhớ lại rằng new(T)phân bổ bộ nhớ bằng 0 cho một mục mới thuộc loại T và trả về địa chỉ của nó.
	Do đó, nó trả về một con trỏ tới loại T. Do đó, arr2thuộc loại *[5]int , ngược lại arr1thuộc loại [5]int .
	Hậu quả là, khi gán một mảng cho một mảng khác,
	một bản sao riêng biệt trong bộ nhớ của mảng được tạo. Ví dụ, khi chúng ta nói:*/
	/*
			TODO:
		arr3 := arr1 - Khai bao arr3 bang arr1
		arr3[2] = 100 - Khi thay doi arr3 thi arr1 khong bi thay doi
		arr3 := &arr1 - Gan nhu nay thi khi thay doi arr3 thi arr1 cung thay doi

	*/
	arr3 := &arr1
	arr3[2] = 100
	/*for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}*/
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%d] = %d\n", i, arr1[i])
	}
	fmt.Printf("================================\n")

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr3[%d] = %d\n", i, arr3[i])

	}
}
