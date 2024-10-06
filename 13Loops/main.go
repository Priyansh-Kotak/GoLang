package main

import "fmt"

func main() {
	fmt.Println("Priyansh is learning Looping")

	// days := []string{"Sun", "Mon", "Tue", "wed", "Thu", "Fri", "Sat"}

	// direct method to print the whole array
	// fmt.Println("Printing days", days)

	// Printing one by one value from the array
	// for i:= 0 ; i< len(days); i++{
	// 	fmt.Println(days[i])
	// }

	//Similar to for each loop but will only get the index fo the elemet from the array
	// for i := range days {
	// 	fmt.Println(i)
	// 	fmt.Println(days[i])
	// }

	// Now we will excess both the index and the value
	// for index, day := range days {
	// 	fmt.Println("index is ", index, " and Value is ", day)
	// }
	// ignoring the index
	// for _, day := range days {
	// 	fmt.Println("index is ignored", " Value is", day)
	// }

	// while loop
	i := 0
	for i < 10 {
		if i == 5 {
			i++
			continue
		}

		if i == 7 {
			goto lco
		}

		fmt.Println(i)
		i++
	}

lco:
	fmt.Println("Redirecting to priyanshkotak.com")

}
