package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2, resultado int
var leyenda string

func main() {
	/* Con fmt */
	fmt.Println("Ingrese numero 1:")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2:")
	fmt.Scanln(&numero2)

	fmt.Printf("Descripcion:")

	/* Con os y bufio */
	scaner := bufio.NewScanner(os.Stdin)
	if scaner.Scan() {
		leyenda = scaner.Text()
	}

	resultado = numero1 + numero2

	fmt.Println(leyenda, resultado)
}
