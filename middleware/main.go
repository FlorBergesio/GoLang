package main

import (
	"fmt"
)

var result int

func main() {
	fmt.Println("Suma de 2 + 3")
	result = operacionesMiddleware(sumar)(2, 3)
	fmt.Println(result)

	fmt.Println("Resta de 7 - 2")
	result = operacionesMiddleware(restar)(7, 2)
	fmt.Println(result)

	fmt.Println("Multiplicacion de 2 * 3")
	result = operacionesMiddleware(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

/* Middleware */
/* Se le pasa por parametro una funcion de un tipo determinado, y devuelve dicha función */
func operacionesMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
