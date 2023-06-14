package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" // string to search
	pat := "[0-9]+.[0-9]+"                                      // pattern search in searchIn

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}
	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchIn, "##.#") // replace pat with "##.#"
	fmt.Println(str)
	// using a function
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

/*
Dòng import "fmt" và import "regexp" là các câu lệnh import để sử dụng package "fmt" và "regexp".

Hàm main là hàm chính của chương trình.

Dòng searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" khai báo một chuỗi để tìm kiếm.

Dòng pat := "[0-9]+.[0-9]+" khai báo biểu thức chính quy để tìm kiếm trong chuỗi searchIn. Biểu thức chính quy này tìm kiếm một dãy số thực trong chuỗi.

Hàm func(s string) string là một hàm được định nghĩa để thực hiện thay thế các mẫu trong chuỗi.

Dòng v, _ := strconv.ParseFloat(s, 32) chuyển đổi giá trị chuỗi s thành một số thực.

Dòng return strconv.FormatFloat(v*2, 'f', 2, 32) chuyển đổi giá trị số thực thành chuỗi và nhân đôi giá trị.

Dòng if ok, _ := regexp.Match(pat, []byte(searchIn)); ok kiểm tra xem chuỗi searchIn có khớp với biểu thức chính quy pat không.

Dòng re, _ := regexp.Compile(pat) biên dịch biểu thức chính quy pat thành một đối tượng regexp để sử dụng cho tìm kiếm và thay thế.

Dòng str := re.ReplaceAllString(searchIn, "##.#") thay thế tất cả các mẫu tìm thấy trong chuỗi searchIn bằng chuỗi "##.#" và gán kết quả cho biến str.

Dòng fmt.Println(str) in ra kết quả của chuỗi sau khi thay thế.

Dòng str2 := re.ReplaceAllStringFunc(searchIn, f) thay thế tất cả các mẫu tìm thấy trong chuỗi searchIn bằng cách sử dụng hàm f và gán kết quả cho biến str2.

Dòng fmt.Println(str2) in ra kết quả của chuỗi sau khi thay thế bằng hàm f.
*/
