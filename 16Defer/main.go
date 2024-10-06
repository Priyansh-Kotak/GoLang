package main

import "fmt"

func main() {
	fmt.Println("Priyansh is learning defer")

	defer fmt.Println("World")
	fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	// defer function will be executed in the reverse order of the declaration

	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("i = ", i)
	}
}
