package main

import "fmt"

/*func f(a [3]int) { fmt.Println(a) } // accepts copy

func fp(a *[3]int) { fmt.Println(a) } // accepts pointer

func main() {
	var ar = [3]int{2, 5, 6}
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}*/

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // dereferencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
