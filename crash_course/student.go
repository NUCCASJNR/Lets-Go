package main

import "fmt"

type Student struct {
	Name         string
	MatricNumber int
	Level        string
	Department   string
}

func StudentSlice() {
	students := []Student{
		{
			Name:         "Rahmon Awwal",
			MatricNumber: 23,
			Level:        "100level",
			Department:   "Mechanical Engineering",
		},
		{
			Name:         "Adegbite Al-Areef",
			MatricNumber: 9,
			Level:        "100level",
			Department:   "Welding and Fabrication Engineering",
		},
		{
			Name:         "Akinola Jamiu",
			MatricNumber: 5,
			Level:        "100level",
			Department:   "Welding and Fabrication Engineering",
		},
	}

	for _, student := range students {
		printStudent(student)
	}
}

func printStudent(s Student) {
	fmt.Println("Name:", s.Name)
	fmt.Println("MatricNumber:", s.MatricNumber)
	fmt.Println("Department:", s.Department)
	fmt.Println("Level:", s.Level)
	fmt.Println("-------------------")
}

