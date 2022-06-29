package main

import (
	"fmt"
	"time"

	us "structs/user"
)

// Herencia
type userAdmin struct {
	us.Usuario
}

func main() {
	User := new(userAdmin)
	User.AltaUsuario(1, "Maria", time.Now(), true)

	fmt.Println(User)
}
