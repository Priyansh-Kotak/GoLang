package main

import "fmt"

// Pt is the Public variable ( and word start with the capital letter will be public)
const Pt = "Priyansh and Teesha"

func main() {
	fmt.Println("# PT teacher forever")
	var username string = "Teesha is a good gir"
	fmt.Println("Hello, " + username + "! Good bye from my life!")
	fmt.Println("Teesha forever")

	var isLoggedIn bool = false
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	fmt.Println(Pt)
}
