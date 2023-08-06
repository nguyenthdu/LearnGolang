package main

import (
	"LearnGolang/Exercise/Employee/department"
	"LearnGolang/Exercise/Employee/employee"
	"fmt"
)

func main() {
	birth1 := "2006-01-02"
	emp1, err := employee.NewEmployee(1, "Nguyen Thanh Du", birth1, 10000)
	if err != nil {
		fmt.Println("Loi nhap ngay", err)
		return
	}
	birth2 := "2002-01-02"
	emp2, err2 := employee.NewEmployee(2, "Nguyen Van A", birth2, 20000)
	if err2 != nil {
		fmt.Println("Loi nhap ngay", err)
		return
	}
	dep := department.NewDepartment(111, "IT")
	dep1 := department.NewDepartment(222, "CC")
	dep.AddEmployee(emp1)
	dep1.AddEmployee(emp2)
	dep.PrintEmployee()
	dep1.PrintEmployee()
}
