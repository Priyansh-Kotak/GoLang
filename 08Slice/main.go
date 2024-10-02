package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Priyansh will learn Slices")

	var st = []string{"Priyansh", "Teesha"}

	fmt.Println("Printing strings", st)

	// st = append(st[1:])
	// fmt.Println("Slicing strings", st)

	st = append(st[:1])
	fmt.Println("Slicing strings", st)

	highScore := make([]int, 4)

	// highScore = append()

	highScore[0] = 222
	highScore[1] = 333
	highScore[2] = 332
	highScore[3] = 331
	// highScore[4] = 338
	// highScore[5] = 111

	// Here we are generating more memory for the new items
	highScore = append(highScore, 338, 111)

	fmt.Println("Printing high score", highScore)

	//isSort the slice
	fmt.Println("isSorted highScore", sort.IntsAreSorted(highScore))

	//Sorting the array slice
	sort.Ints(highScore)
	fmt.Println("Sorting highScore", highScore)

	//isSort the slice
	fmt.Println("isSorted highScore", sort.IntsAreSorted(highScore))

	// removing the 3 item from the array slice

	var sports = []string{"volly", "football", "cricket", "basketball", "skettings"}
	fmt.Println("Pringting out Sports array:=", sports)

	var index int = 2
	sports = append(sports[:index], sports[index+1:]...)
	fmt.Println("Removed the third item from the sport array ", sports)

}
