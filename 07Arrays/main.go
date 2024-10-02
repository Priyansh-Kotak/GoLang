package main

import "fmt"

func main() {
	fmt.Println("Priyansh will learn arrays")

	var num [5]int

	num[0] = 1
	num[1] = 2
	num[2] = 3

	fmt.Println("Printing array", num)

	// Implicit declaratoin of the array
	var myString = [3]string{"Priyash", "Teesha"}
	fmt.Println("Printing string array", myString)

}
