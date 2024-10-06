package main

import "fmt"

func main() {
	fmt.Println("Priyansh will lear the structre because go lang don't have classes")

	priyansh := User{"Priyansh", 22, "priyanshkotak@gmail.com", "9978402359"}

	fmt.Println(priyansh)
	fmt.Printf("Printing the structure value %+v\n", priyansh) // this is used to print the full struct with it's values
	fmt.Println(priyansh.Age)  // this will simply print the age of the struct

}

// We are using Capital letter in the first place of the word the reason is that we need to export it to another package
type User struct {
	Name         string
	Age          int
	Email        string
	Phone_Number string
}
