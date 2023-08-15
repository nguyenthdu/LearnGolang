package main

import "fmt"

func main() {
	// initialization within switch block
	switch num1 := 100; {
	case num1 < 0:
		fmt.Println("Number is negative")

	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")

	default:
		fmt.Println("Number is 10 or greater")
	}
}

/*switch initialization; {
case val1:
...
case val2:
...
default:
...
}*/
/*
switch a, b := x[i], y[j]; {
case a < b: t = -1
case a == b: t = 0
case a > b: t = 1
}*/
