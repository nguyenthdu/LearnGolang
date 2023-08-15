package main

import (
	"bytes"
	"fmt"
)

/*
import "bytes"

	type Buffer struct {
	  ...
	}
*/
//create buffer
var buffer bytes.Buffer

// create point buffer
var r *bytes.Buffer = new(bytes.Buffer)

// create func
//func NewBuffer(buf []byte) *Buffer

func main() {
	var b bytes.Buffer
	// Write strings to the Buffer.

	b.WriteString("ABC")
	b.WriteString("DEF")
	// Use Fprintf with Buffer.
	fmt.Fprintf(&b, " A number: %d, a string: %v\n", 10, "bird")
	b.WriteString("[DONE]")
	// Convert to a string and print it.
	fmt.Println(b.String())
}

// TODO: Noi chuoi
/*Phương pháp này tiết kiệm bộ nhớ và CPU hơn nhiều so với += , đặc biệt nếu số lượng chuỗi để nối lớn.
func noiChuoi() {
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok { //method getNextString() not shown here
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")
}
*/
