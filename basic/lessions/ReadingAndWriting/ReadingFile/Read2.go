package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	inputFile := "D:/LearnGolang/basic/lessions/ReadingAndWriting/ReadingFile/product.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", string(buf))
	os.Stdout.Write(buf)
}
