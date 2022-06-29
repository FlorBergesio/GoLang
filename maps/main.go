package main

import "fmt"

func main() {
	paises := make(map[string]string)
	paises["Argentina"] = "Catamarca"
	fmt.Println(paises)

	edades := map[string]int{
		"Maria":    17,
		"Veronica": 28,
		"Pablo":    40,
	}
	fmt.Println(edades)

	// Agregar un elemento
	edades["Emanuel"] = 30

	// Borrar un elemento
	delete(edades, "Pablo")

	fmt.Println(edades)

	// Recorrer el map
	for persona, edad := range edades {
		fmt.Printf("La persona %s tiene %d años", persona, edad)
	}

	// Conseguir el valor de un elemento y, a la vez, averiguar si existe
	edad, existe := edades["Gabriela"]
	fmt.Printf("La edad capturada es de: %d; La persona existe? %t", edad, existe)
}
