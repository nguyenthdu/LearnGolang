package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	CopyFile("D:/LearnGolang/basic/lessions/ReadingAndWriting/target.txt", "D:/LearnGolang/basic/lessions/ReadingAndWriting/source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
