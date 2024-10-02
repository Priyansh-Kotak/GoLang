package main

import "fmt"

func main() {
	fmt.Println("Priyansh is learning function")

	// we cannot call function inside the function
	// func greeterTwochan()  {
	// 	fmt.Println("We are inside th greeter Two");
	// }

	// greeterTwochan()

	greeter()

	fmt.Println("Parameteraized function where there are arguments ")

	// taking the user input

	var num1 int
	var num2 int

	fmt.Println("Enter number 1")
	_, err := fmt.Scan(&num1)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	fmt.Println("Enter number 2")
	// _, err := fmt.Scan(&num2)
	_, erro := fmt.Scan(&num2)
	if erro != nil {
		fmt.Println("Error", erro)
	}

	adder(num1, num2)

	// fmt.Println("result", result)

	prores, myString := multipleAdder(1, 2, 3, 4)

	fmt.Println("My pros", prores)
	fmt.Println("My pros", myString)

}

func greeter() {
	fmt.Println("Hello ji Priyansh this side")

	// annonymisous function are only called inside another function
	func() {
		fmt.Println("Using annonymisous function")
	}()

}

func adder(val1 int, val2 int) {
	add := val1 + val2

	fmt.Println("Addition is", add)

}

// variadic functions means functions which takes multiple arguments and returns something
func multipleAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "I am returning the value"
}
