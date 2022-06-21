package main

import "fmt"

// Creamos función Tabla, que a su vez devuelve una funcion que devuelve un entero
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

func main() {
	tablaDe := 3
	tablaClosure := Tabla(tablaDe)

	for i := 1; i < 11; i++ {
		fmt.Println(tablaClosure())
	}
}
