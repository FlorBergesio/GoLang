package main

import (
	"fmt"
)

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 1 + 2 = %d", Calculo(1, 2))

	// Redefinimos Calculo
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resta 8 - 2 = %d", Calculo(8, 2))
}
