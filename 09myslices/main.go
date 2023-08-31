package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices golang")

	var fruitList = []string{"Apple", "Banana", "Peach"}

	fmt.Println("Fruitlist is: ", fruitList)

	fruitList = append(fruitList, "Mango", "Jamfruit")
	fmt.Println("Fruitlist is: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Fruitlist is : ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 345
	highScores[1] = 212
	highScores[2] = 576
	highScores[3] = 321
	// highScores[4] = 435

	fmt.Println("highscores is : ", highScores)

	highScores = append(highScores, 987, 232, 676)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted highscores : ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
