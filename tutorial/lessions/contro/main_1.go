package main

import "fmt"

type Person struct {
	name string
	age  int
}

func updateAge(ptr *Person, newAge int) {
	ptr.age = newAge

}
func updateAgeNo(ptr Person, newAge int) Person {
	ptr.age = newAge
	return ptr

}

func main() {
	//var ptr *int
	//var x int = 10
	//ptr = &x // lay dia chi bien x gan cho con tro ptr
	//var n int = 7
	//ptr = &n
	//println("Value of x is: ", x)
	//println("Value of x is: ", *ptr)
	//var z = 10
	//fmt.Println("Value of z is: ", z)
	//changeValue(&z)
	//fmt.Println("Value of z is: ", z)
	p := Person{name: "John", age: 21}
	fmt.Println("Person before update: ", p)
	//updateAge(&p, 22)
	p = updateAgeNo(p, 22)
	fmt.Println("Person after update: ", p)
}

func changeValue(ptr *int) {
	*ptr = 20
}
