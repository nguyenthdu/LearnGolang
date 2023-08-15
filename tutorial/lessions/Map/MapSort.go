package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key: %v, value: %v / ", k, v) // read random keys
	}
	keys := make([]string, len(barVal)) // storing all keys in separate slice
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys) // sorting the keys slice
	fmt.Println()
	fmt.Println("\nsorted:")
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v / ", k, barVal[k]) // reading key from keys and value from barVal1
	}
}
