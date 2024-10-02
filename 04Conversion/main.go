package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Println("My name is ", name)

	numRating, err := strconv.ParseInt(strings.TrimSpace(name), 10, 64)

	if err != nil {
		fmt.Println("Error in converting string to float", err)
	} else {
		fmt.Println("The number you entered is", numRating+1)
	}
}
