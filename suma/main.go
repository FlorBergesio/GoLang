package main

import (
	"fmt"
)

func main() {
	fmt.Println(suma(1, 2, 3, 4))
	fmt.Println(suma(7, 245, 3, 4))
	fmt.Println(suma(1, 2))
	fmt.Println(suma(4))
}

func suma(numero ...int) (total int) {
	total = 0
	for _, num := range numero {
		total = total + num
	}
	return
}
