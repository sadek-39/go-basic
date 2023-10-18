package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	sadek := User{"Sadek", "sadeksltn@gmail.com", true, 28}

	fmt.Println(sadek)
	sadek.GetStatus()
	sadek.SetEmail("test@gmail.com")
	fmt.Println(sadek.Email)

}

func (u User) GetStatus() {
	fmt.Println("The user is active: ", u.Status)
}

func (u User) SetEmail(email string) {
	u.Email = email
	fmt.Println(u.Email)
}
