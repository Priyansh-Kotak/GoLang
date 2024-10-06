package main

import (
	"fmt"
)

func main() {

	fmt.Println("Priyansh Kotak")

	// reader := bufio.NewReader(os.Stdin)

	var num int

	fmt.Println("Enter the number : - ")

	_, err := fmt.Scan(&num)

	var result string

	if num > 10 {
		result = "Less than 10"
	} else if num < 10 {
		result = "Greater than 10"
	} else {
		result = "Equal to 10"
	}

	fmt.Println("Result is :-", result)

	if 9%2 == 0 {
		fmt.Println("Result is even")
	} else {
		fmt.Println("Resust is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Number is :-", num)
	}

}
