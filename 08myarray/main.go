package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "PineApple"
	fruitList[3] = "Peach"

	fmt.Println("Fruitlist is: ", fruitList)
	fmt.Println("Length of fruitlist is: ", len(fruitList))

	var vegList = [5]string{"Spinach", "Raddish", "Karibevu"}

	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Length of veg list: ", len(vegList))

}
