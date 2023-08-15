package main

import (
	"fmt"
)

var (
	barVal1 = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43,
		"lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal1)) // interchanging types of keys and values
	for k, v := range barVal1 {
		invMap[v] = k // key becomes value and value becomes key
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("key: %v, value: %v / ", k, v)
	}
}
