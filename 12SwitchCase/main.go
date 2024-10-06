package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Priyansh is learning Switch Case")

	// rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())

	rng := rand.New(source)

	randomNumber := rng.Intn(6) + 1 // generate random number b/w 0 to 6

	fmt.Println(randomNumber)

	switch randomNumber {
	case 1:
		fmt.Println("Dice number is 1")

	case 2:
		fmt.Println("Dice number is 2")
	case 3:
		fmt.Println("Dice number is 3")
	case 4:
		fmt.Println("Dice number is 4")
		fallthrough // continue to next case
	case 5:
		fmt.Println("Dice number is 5")
	case 6:
		fmt.Println("Dice number is 6")

	}

}
