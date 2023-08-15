package main

import (
	"errors"
	"fmt"
)

/*
	if value, err := pack1.Func1(param1); err != nil {
	  fmt.Printf("Error %s in pack1.Func1 with parameter %v", err.Error(), param1)
	  return // or: return err
	} else {

	   // process(value)
	}
*/
//Interface error
/*Các giá trị lỗi được sử dụng để chỉ ra trạng thái bất thường.
Gói errorschứa một errorStringcấu trúc, thực hiện giao diện lỗi.
Để dừng thực thi chương trình ở trạng thái lỗi, chúng ta có thể sử dụng os.Exit(1).*/
/*type error interface {
	Error() string
}
*/
/*error interface
errors.New(text string) error
err.Error() string*/
//TODO: Defining error
/*Bất cứ khi nào bạn cần một loại lỗi mới, bạn có thể tạo một loại có chức năng
errors.Newtừ errors gói (mà bạn sẽ phải nhập)
và cung cấp cho nó một chuỗi lỗi thích hợp, như sau:*/
//err := errors.New("math - square root of negative number")

var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}

// error: Not found error
// hàm kiểm tra tham số của hàm căn bậc hai
/*func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	// implementation of Sqrt
}
//Bạn có thể gọi chức năng này như sau:
if f, err := Sqrt(-1); err != nil {
  fmt.Printf("Error: %s\n", err)
}
*/
