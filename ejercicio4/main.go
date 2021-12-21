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

type Tienda struct {
	Users    []User
	Products []Item
}

func main() {
	product := new(Item)
	newProduct := product.addItem("lata", 1200.00, 10)
	tienda := new(Tienda)
	tienda.Products = append(tienda.Products, newProduct)
	user := User{Name: "Laynerker", LastName: "Guerrero", Email: "Laynerker.gdl@gmail.com", Items: []Item{}}
	tienda.Users = append(tienda.Users, user)
	tienda.addItemUser("Laynerker", newProduct)
	fmt.Println("esta de tienda: ", tienda)
	tienda.deleteProducts("Laynerker")
	fmt.Println("esta de tienda: ", tienda)
}

func (i Item) addItem(name string, price float64, count int) Item {
	i.name = name
	i.price = price
	i.count = count
	fmt.Println("Se agrego el producto: ", i)
	return i
}

func (t *Tienda) addItemUser(nameUser string, item Item) {
	for k, u := range t.Users {
		if u.Name == nameUser {
			t.Users[k].Items = append(t.Users[k].Items, item)
		}
	}
	fmt.Println("Se agrego al usuario ", nameUser, "el producto ", item)

}

func (t *Tienda) deleteProducts(nameUser string) {
	for k, u := range t.Users {
		if u.Name == nameUser {
			t.Users[k].Items = []Item{}
		}
	}
	fmt.Println("Se eliminaron los productos del usuario", nameUser)
}
