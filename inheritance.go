// Giả sử ta có một hệ thống quản lý nhân viên, có nhân viên chung (Employee) và nhân viên quản lý (Manager):
package main

import "fmt"

// Struct Employee
type Employee struct {
	Name string
}

func (e Employee) Work() {
	fmt.Println(e.Name, "đang làm việc")
}

// Struct Manager kế thừa từ Employee
type Manager struct {
	Employee // Embedding Employee
	TeamSize int
}

func (m Manager) Lead() {
	fmt.Println(m.Name, "đang lãnh đạo đội nhóm có", m.TeamSize, "người")
}

func main() {
	emp := Employee{Name: "Nguyễn Văn B"}
	emp.Work() // Nguyễn Văn B đang làm việc

	mgr := Manager{Employee: Employee{Name: "Trần Văn C"}, TeamSize: 5}
	mgr.Work() // Trần Văn C đang làm việc (Kế thừa từ Employee)
	mgr.Lead() // Trần Văn C đang lãnh đạo đội nhóm có 5 người
}
