package employee

import "time"

type Employee struct {
	id     int
	name   string
	birth  time.Time
	salary float64
}

func NewEmployee(id int, name string, birth string, salary float64) (*Employee, error) {
	// định dạng birth theo kiểu thời gian
	

	t, err := time.Parse("2006-01-02", birth)
	if err != nil {
		return nil, err
	}
	return &Employee{
		id:     id,
		name:   name,
		birth:  t,
		salary: salary,
	}, nil
}

// tìm nhân viên theo id
func (emp *Employee) FindEmployeeById(id int) *Employee {
	if emp.id == id {
		return emp
	}
	return nil
}

// in ra thông tin nhân viên
func (emp *Employee) PrintEmployee() *Employee {
	return emp
}
