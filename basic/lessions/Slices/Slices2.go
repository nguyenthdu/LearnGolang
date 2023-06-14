package main

import "fmt"

func main() {
	/*sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)                    // output: [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) // n == 3
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3) // output: [1 2 3 4 5 6]*/
	/*var t int
	s := []int{1, 2, 3}
	fmt.Println("Nhap vao chieu dai cua mang: ")
	fmt.Scan(&t)
	e := enlarge(s, t)
	fmt.Printf("Chieu dai cua mang sau khi da them la: %d", len(e))
	fmt.Print(e)*/
	slice := []string{"M", "N", "O", "P", "Q", "R"}
	insert := []string{"A", "B", "C"}
	a := insertSlice(slice, insert, 0)
	fmt.Println(a)
}
func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	return ns
}
func insertSlice(slice, insertion []string, index int) []string {
	slice = append(slice[:index], append(insertion, slice[index:]...)...)
	return slice
}

/*
func insertSlice(slice, insertion []string, index int) []string {
    result := make([]string, len(slice) + len(insertion))
	at := copy(result, slice[:index])// sao chep tu day den index cua slice vao result va luu cac gia tri va at
	at += copy(result[at:], insertion) - noi dung tu insertion sao chep vao result phai sau slice
	Câu lệnh at += copy(result[at:], insertion)sẽ đặt phần tử insertion  sau phần tử slice trong result và thêm phần tử đã sao chép vào phần tử at.
    copy(result[at:], slice[index:])
    return result
}
*/

func Filter(s []int, fn func(int) bool) []int {
	var t []int
	for _, v := range s {
		if even(v) { //s[i] cung duoc
			t = append(t, v)
		}
	}
	return t
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
