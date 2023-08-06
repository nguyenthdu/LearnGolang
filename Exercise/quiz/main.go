package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	//_ = csvFilename
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		//os.Exit(1)
	}
	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provide CSV file.")
	}
	//_ = file
	fmt.Println(line)
}

// TODO: STRUCT SAVE QUESTION AND ANSWER
type problem struct {
	q string // question
	a string // answer
}

// TODO: FUNCTION PARSE LINES
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
