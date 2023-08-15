package main

import "fmt"

// Tao 1 slices: var slice1 []type = make([]type, len) hoac slice1 := make([]type, len)
/*
make([]int, 50, 100)
new([100]int)[0:50]
*/
func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!
	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice:
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	// grow the slice beyond capacity:
	// . slice1 = slice1[0:7] // panic: runtime error: slice bounds out of range

}

/*func main() {
  var slice1 []int = make([]int, 10)
  // load the array/slice:
  for i := 0; i < len(slice1); i++ {
    slice1[i] = 5 * i
  }
  // print the slice:
  for i := 0; i < len(slice1); i++ {
    fmt.Printf("Slice at %d is %d\n", i, slice1[i])
  }
  fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
  fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}*/
