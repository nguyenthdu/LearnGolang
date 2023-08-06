package department

import (
	"LearnGolang/Exercise/Employee/employee"
	"fmt"
)

type Department struct {
	id       int
	name     string
	employee []*employee.Employee
}

func NewDepartment(id int, name string) *Department {
	return &Department{
		id:   id,
		name: name,
	}
}
func (dept *Department) AddEmployee(emp *employee.Employee) {
	dept.employee = append(dept.employee, emp)
}

// in ra danh sách nhân viên của phòng ban
func (dept *Department) PrintEmployee() {
	//duyệt danh sách department
	fmt.Println("Department:", dept.id)
	fmt.Println("Department:", dept.name)
	for _, emp := range dept.employee {
		fmt.Println("Employee:", emp.PrintEmployee())
	}

}

// tìm phòng ban theo id và in ra nhân viên trong phòng ban đó
func (dept *Department) FindDepartmentById(id int) *Department {
	if dept.id == id {
		return dept
	}
	return nil
}
