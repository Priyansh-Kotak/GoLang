package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("The choice is yours")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter how much you love Teesha ?")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	fmt.Printf("Type of your rating %T", input)

	// Simple user input

	var num int

	fmt.Println("Enter the number : - ")

	_, err := fmt.Scan(&num)

	// var result string

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Number is :-", num)
	}

}
