package main

import (
	"fmt"
	"os"
)

func main() {
	datos := []byte("ID;Precion;cantidad\n1;30012.00;2\n2;1000000.00;1\n3;50.50;1")
	err := os.WriteFile("../productos.csv", datos, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Todo bien")
}
