package main

import "fmt"

func maps() {
	users := map[string]any{
		"Name": "Adegbite Al-Areef",
		"Occupation": "Welding intern/Backend engineer",
		"Age": 20,
	}
	fmt.Println(users)
}