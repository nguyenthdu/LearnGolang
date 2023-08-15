package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// Một ParseError chỉ ra một lỗi trong việc chuyển đổi một từ thành một số nguyên.
type ParseError struct {
	Index int    // Chỉ mục vào danh sách các từ được phân tách bằng dấu cách.
	Word  string // Từ tạo ra lỗi phân tích cú pháp.
	Err   error  // Lỗi thô gây ra lỗi này, nếu có.
}

// Chuỗi trả về thông báo lỗi mà con người có thể đọc được.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parse %q as int", e.Word)

}

// Phân tích cú pháp phân tích các từ được phân tách bằng dấu cách trong đặt dưới dạng số nguyên.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg:%v", r)
			}
		}
	}()
	fileds := strings.Fields(input)
	numbers = fileds2numbers(fileds)
	return

}
func fileds2numbers(fileds []string) (numbers []int) {
	if len(fileds) == 0 {
		panic("no word  to parse")
	}
	for idx, field := range fileds {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
