package main

import "fmt"

var Days = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
} // do initialization here

func findDay(n int) string {
	val, isPresent := (Days[n])
	if isPresent { // what if key is not present
		return val
	} else {
		return "Wrong Key"
	} // return wrong key if invalid key
}
func main() {
	a := findDay(20)
	fmt.Println(a)
}
