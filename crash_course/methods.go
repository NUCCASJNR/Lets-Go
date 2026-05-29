package main

import "fmt"

type Lecturer struct {
	Name       string
	Department string
}

func (l Lecturer) introduce() {
	fmt.Println("My name is", l.Name)
	fmt.Println("Department:", l.Department)
}

func PrintLecturer()  {
	l := Lecturer{
		Name: "Adewale Samuel",
		Department: "Welding and fabrication engineering",
	}
	l.introduce()
}
