package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary int
}
type Skills []string

// Embedding struct in struct
type People struct {
	name string
	age  int
}
type Staff struct {
	People
	salary1 float32
}
type Manager struct {
	People
	catory string
	Skills
}

func main() {
	// cach khi bao truc go
	var emp Employee
	emp.name = "Nguyen Thanh Du"
	emp.age = 12
	emp.salary = 50000
	fmt.Println(emp)
	// cach 2
	emp1 := Employee{name: "Nguyen Van A", age: 44, salary: 33}
	fmt.Println(emp1)
	//cach 3
	emp2 := Employee{"Nguyen Van c", 23, 44}
	fmt.Println(emp2)
	//struct anonymous
	emp4 := struct {
		name   string
		age    int
		salary int
	}{"Nguyen Van T", 34, 33}
	fmt.Println(emp4, "\n========================================")
	// su dung struct in struct
	var staff Staff
	staff.name = "Thanh"
	staff.age = 33
	staff.salary1 = 123
	fmt.Println(staff)
	staff1 := Staff{People: People{name: "thanh thanh", age: 33}, salary1: 1111}

	fmt.Println(staff1.People.name, staff1.age, staff1.salary1)
	man := Manager{People: People{name: "QQQQQQ", age: 333}, catory: "PM", Skills: []string{"a", "b"}}
	fmt.Println(man)

}
