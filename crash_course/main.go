package main

import (
	"fmt"
)

func new() {
    name := "Areef"
    age := 18
    occupation := "Intern"

    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(occupation)
}


func greet(name string) string {
    return "Hello " + name + " how are you doing"
}
func add(a int, b int) int {
    return a + b
}

func main() {
    // message := greet("Areef")
    // result := add(9, 5)
    // new()
    // printDetails()
    // statement()
    // fmt.Println(message)
    // fmt.Println(result)
    // slices()
    // maps()
    // combined()
    // result()
    // StudentSlice()
    // PrintLecturer()
    // fmt.Println("=== NEW MAIN RUNNING ===")
    // NewPrint()
    // PrintStudentManagement()
    PrintStudentApi()
}