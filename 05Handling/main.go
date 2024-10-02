package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Priyansh will learn the time in go lang")

	presentTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Hour(), time.Now().Nanosecond(), time.Now().Location())

	fmt.Println(presentTime)
}
