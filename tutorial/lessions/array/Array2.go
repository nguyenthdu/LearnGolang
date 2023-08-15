package main

import "fmt"

func main() {
	//duyetmang()
	//themDau()
	//themCuoi()
	//themGiua()
	//xoaDau()
	//xoaCuoi()
	//xoaGiua()
	//capNhat()
	//noiMang()

}

// TODO: duyet mang
func duyetmang() {
	numbers := [5]int{10, 20, 30, 40, 50}
	//leght := len(numbers) lay do  dai mang
	//numbers = append(numbers, 4, 5) them 1 hoac nhieu phan tu vao cuoi mang tra ve mang
	//subArray := numbers[1:4] - cat: truy cap phan tu cua mang chi dinh pham vi
	/* copy
	source := []int{1, 2, 3}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Println(destination) // Kết quả: [1 2 3]

	*/
	//Su dung for
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Print("=======================================")
	//Su dung range
	for index, values := range numbers {
		fmt.Println(index, ":", values)
	}
}

// TODO: them phan tu vao mang
func themDau() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers = append([]int{60}, numbers...)
	fmt.Println(numbers) // Kết quả: [60 10 20 30 40 50]

}

// TODO: them vao cuoi
func themCuoi() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers = append(numbers, 100)
	fmt.Println(numbers)
}

// TODO: add emlement in center array
func themGiua() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers = append(numbers[:2], append([]int{77}, numbers[2:]...)...)
	fmt.Println(numbers) // Kết quả: [10 20 77 30 40 50]
}

// TODO: xoa phan tu  dau mang
func xoaDau() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers = numbers[1:]
	fmt.Println(numbers)
}

// TODO: Xoa phan tu cuoi mang
func xoaCuoi() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers = numbers[:len(numbers)-1]
	fmt.Println(numbers)
}

// TODO: Xoa phan tu giua mang
func xoaGiua() {
	numbers := []int{10, 20, 30, 40, 50}
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println(numbers)
}

// TODO: Cap nhat phan tu trong mang
func capNhat() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers[1] = 11
	fmt.Println(numbers)
}
func noiMang() {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	numbers3 := append(numbers1, numbers2...)
	fmt.Println(numbers3) // Kết quả: [1 2 3 4 5 6]
}
func xoaMang() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = []int{}
	fmt.Println(numbers) // Kết quả: []
}
func xoaNhieu() {
	numbers := []int{1, 2, 3, 4, 5}
	index := 1
	count := 2
	numbers = append(numbers[:index], numbers[index+count:]...)
	fmt.Println(numbers) // Kết quả: [1 4 5]
}
