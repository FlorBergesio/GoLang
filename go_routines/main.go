package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go name_slow("Antonella")
	fmt.Println("Hola Mundo")
	var x string
	fmt.Scanln(&x)
}

func name_slow(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
