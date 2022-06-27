package main

import "fmt"

var array1 [10]int
var matriz1 [5][7]int

func main() {
	array1[0] = 4
	array1[5] = 40

	array2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice := []int{1, 2, 3}

	fmt.Println(array1)

	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	matriz1[3][4] = 1

	fmt.Println(matriz1)

	fmt.Println(slice)

	variante()
	variante2()
}

func variante() {
	elementos := [5]int{1, 2, 3, 4, 5}
	// crear un slice de los elementos desde la posición 3 en adelante:
	porcionFinal := elementos[3:]
	fmt.Println(porcionFinal)

	// crear un slice de los elementos desde el primer elemento hasta la posición 3:
	porcionInicial := elementos[:3]
	fmt.Println(porcionInicial)
}

func variante2() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad max: %d", len(elementos), cap(elementos))
}
