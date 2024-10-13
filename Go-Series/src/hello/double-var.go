package main

import "fmt"

func Double() {
	var age int // variable declaration
	fmt.Println("My initial age is", age)
	age = 29 //assignment
	fmt.Println("My age after first assignment is", age)
	age = 54 //assignment
	fmt.Println("My age after second assignment is", age)
}



func Tripple() {
	var name string
	name = "Al-areef"
	fmt.Println("My FirstName name is: ", name)
	name = "Adegbite"
	fmt.Println("My Surname is: ", name)
	name = "Ayomipo"
	fmt.Println("My lastname is: ", name)
}
func main() {
	Double()
	Tripple()
}