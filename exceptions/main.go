package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "file.txt"
	file, err := os.Open(fileName)

	defer file.Close()

	ejemploPanic()

	if err != nil {
		fmt.Println("Error abriendo el programa")
		os.Exit(1)
	}
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Hubo un panic \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("Error. A es igual a 1")
	}
}
