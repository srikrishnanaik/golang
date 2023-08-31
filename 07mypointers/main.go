package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int

	// fmt.Println("Value of pointer is ", ptr)

	mynumber := 23

	var ptr = &mynumber

	fmt.Println("Address of mynumber", ptr)
	fmt.Println("Value of mynumber", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value: ", mynumber)

}
