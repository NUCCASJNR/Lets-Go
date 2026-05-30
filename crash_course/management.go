package main

import "fmt"

type StudentManagement struct {
	Name         string
	MatricNumber int
	Level        int
	Department   string
}

func (s StudentManagement) Introduce() {
	fmt.Println("Hello, my name is", s.Name)
	fmt.Println("Department:", s.Department)
	fmt.Println("Level:", s.Level)
	fmt.Println("-------------------")
}

func (s *StudentManagement) Promote() {
	s.Level += 100
}

func (s *StudentManagement) ChangeDepartment(newDept string) {
	s.Department = newDept
}

func PrintStudentManagement() {
	students := []StudentManagement{
		{
			Name:         "John Doe",
			MatricNumber: 10,
			Level:        100,
			Department:   "Accounting",
		},
		{
			Name:         "Allison Becker",
			MatricNumber: 11,
			Level:        200,
			Department:   "Physics",
		},
		{
			Name:         "Cole Palmer",
			MatricNumber: 12,
			Level:        300,
			Department:   "Attacking Midfielder",
		},
	}

	for i := range students {
		fmt.Println("BEFORE UPDATE:")
		students[i].Introduce()

		students[i].Promote()
		students[i].ChangeDepartment("Mechanical Engineering")

		fmt.Println("AFTER UPDATE:")
		students[i].Introduce()
	}
}
