package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D:/LearnGolang/basic/lessions/ReadingAndWriting/ReadingFile/product2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3) // scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

/*
import "path/filepath"
import "fmt"

func main() {
    path := "/usr/local/bin/myfile.txt"
    base := filepath.Base(path)
    fmt.Println(base)
}
Kết quả sẽ là "myfile.txt". Hàm Base() đã trả về phần tử cuối cùng của đường dẫn, tức là tên tập tin "myfile.txt", mà không bao gồm ký tự ngăn cách cuối cùng ("/").
*/
