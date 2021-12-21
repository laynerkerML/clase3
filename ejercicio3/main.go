package main

import (
	"fmt"
)

type User struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func main() {
	user := new(User)
	fmt.Println("base inicio al", user)
	user.changeDatosS("name", "Laynerker")
	user.changeDatosS("last_name", "Laynerker")
	user.changeDatosS("email", "Laynerker")
	user.changeDatosS("password", "Laynerker")
	user.changeDatosI(31)
	fmt.Println("user final: ", user)
}

func (u *User) changeDatosS(field string, value string) {
	switch field {
	case "name":
		u.Name = value
	case "last_name":
		u.LastName = value
	case "email":
		u.Email = value
	case "password":
		u.Password = value
	}
}

func (u *User) changeDatosI(value int) {
	u.Age = value
}
