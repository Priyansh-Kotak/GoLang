package main

import "fmt"

func main() {
	fmt.Println("Priyansh will learn the pointer and the address concepts")

	// var ptr *int

	// fmt.Println("ptr valur ", ptr)

	number := 1000

	var ptr = &number

	fmt.Println("Address of the pointer ", ptr)

	*ptr = *ptr + 1

	fmt.Println("Value of the pointer ", *ptr)
}
