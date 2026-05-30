package main

import "fmt"

type NewStudent struct {
	Name         string
	MatricNumber int
	Level        int
	Department   string
}

// METHOD (belongs to Student)
func (s NewStudent) printInfo() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Matric:", s.MatricNumber)
	fmt.Println("Level:", s.Level)
	fmt.Println("Department:", s.Department)
	fmt.Println("-------------------")
}

// POINTER METHOD (modifies struct)
func (s *NewStudent) updateDepartment(newDept string) {
	s.Department = newDept
}

// func (s *NewStudent) Promote()

func NewPrint() {
	s := NewStudent{
		Name:         "Areef",
		MatricNumber: 9,
		Level:        100,
		Department:   "Welding",
	}
	// s.Promote()
	s.printInfo()

	s.updateDepartment("Mechanical Engineering")

	fmt.Println("After update:")
	s.printInfo()
}