package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	datos, err := os.ReadFile("../productos.csv")
	if err != nil {
		fmt.Println("err: ", err)
	}
	datosStr := string(datos)
	lineas := strings.Split(datosStr, "\n")
	for _, line := range lineas {
		columna := strings.Split(line, ";")
		/*for _, colum := range columna {
			fmt.Print(colum, "\t")
		}*/
		res := fmt.Sprintf("%s\t\t%s\t\t%s", columna[0], columna[1], columna[2])
		fmt.Println(res)
	}
	fmt.Println(len(lineas))
}
