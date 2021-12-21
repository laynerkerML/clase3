package main

import (
	"fmt"
)

type Item struct {
	name  string
	price float64
	count int
}

type User struct {
	Name     string
	LastName string
	Email    string
	Items    []Item
}

func main() {
	fmt.Println("hola")
}

func (i Item) addItem(name string, price float64, count int) Item {
	i.name = name
	i.price = price
	i.count = count
	return i
}

func (u *User) addItemUser(item Item) {
	u.Items = append(u.Items, item)
}

func (u *User) deleteProduct(name User) {

}
