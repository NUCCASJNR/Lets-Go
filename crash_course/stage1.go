package main

import "fmt"

func stageOne(name string, age int, isStudent bool) {
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	fmt.Println("Name:", name)
	fmt.Println("Student:", isStudent)
}

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func sliceDemo() {
	numbers := []int{5, 10, 15}
	numbers = append(numbers, 20)

	fmt.Println(numbers)
}

func mapUser() {
	users := map[string]any{
		"Name":       "Adegbite Al-Areef Ayomipo",
		"Department": "Welding and Fabrication Engineering",
		"School":     "Yaba College of Technology",
	}

	fmt.Println(users)
}

func combined() {
	
	stageOne("Al-Areef", 18, true)
	printNumbers()
	sliceDemo()
	mapUser()
}
