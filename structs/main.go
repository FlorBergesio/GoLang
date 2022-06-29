package main

import (
	"fmt"
	"time"
)

// Definición de estructura - equivalente a definir una clase
type usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func main() {
	User := new(usuario)
	User.Id = 10
	User.Nombre = "Maria"

	fmt.Println(User)
}
