package main

import (
	"fmt"
)

func main() {
	fmt.Println("Priyansh is learning the map")

	langauge := make(map[string]string)

	langauge["Js"] = "Javascript"
	langauge["Go"] = "GoLang"
	langauge["Java"] = "Java"
	langauge["Python"] = "Python"

	fmt.Println("Printing the map", langauge)
	fmt.Println("Individual printing", langauge["Js"])

	delete(langauge, "Js")
	fmt.Println("Printing map after deleting", langauge)

	//Looping the map

	for key, value := range langauge {
		fmt.Println("Key :- ", key, ", Value :-", value, "\n")
	}

}
