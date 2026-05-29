package main

import "fmt"

type User struct {
	Name       string
	Age        int
	Department string
	IsStudent  bool
}

func checkUser(user User) {
	if user.Age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Student:", user.IsStudent)
	fmt.Println("Department:", user.Department)
	fmt.Println("Age:", user.Age)
	fmt.Println("-------------------")
}

func result() {
	users := []User{
		{
			Name:       "Al-Areef",
			Age:        20,
			Department: "Welding",
			IsStudent:  true,
		},
		{
			Name:       "John",
			Age:        20,
			Department: "Welding",
			IsStudent:  true,
		},
	}

	for _, user := range users {
		checkUser(user)
	}
}

