package main

import "fmt"

func main() {
	fmt.Println("Priyansh is learning methods")

	priyansh := User{"Priyansh", "priyanshkotak1@gmail.com", true, 19, false}

	fmt.Println(priyansh)
	priyansh.GetStatus()
	priyansh.NewEmail()
}

type User struct {
	Name      string
	Email     string
	Status    bool
	Age       int
	ismarried bool
}

// we will methods
func (u User) GetStatus() {
	u.Status = false
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "hardkotak@gmail.com"
	fmt.Println("New email is :", u.Email)
}
