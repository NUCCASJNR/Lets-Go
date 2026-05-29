package main

import "fmt"

func introduce(name, department, school string) {
    fmt.Println("My name is " + name)
    fmt.Println("A student in the department of: " + department)
    fmt.Println("In: " + school)
}

func printDetails() {
    introduce("Areef", "Welding & Fabrication Engineering", "YABATECH")
}