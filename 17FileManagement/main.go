package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Starting file management service...")

	content := "My name is Priyansh and me and Teesha are forever"

	file, err := os.Create("./test_file_management_service")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of the file is := ", length)

	defer file.Close()

	readFile("./test_file_management_service")
}

func readFile(path string) {
	databyte, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data from the file is :=", string(databyte))
}
