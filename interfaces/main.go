package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	genero() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

/* Before refactoring */
/*
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}
*/

/* After refactoring */
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) genero() string {
	if h.esHombre == true {
		return "Masculino"
	} else {
		return "Femenino"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Mi género es: %s; y estoy respirando\n", hu.genero())
}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) esCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalRespirando(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalCarnivoro(an animal) int {
	if an.esCarnivoro() == true {
		return 1
	}
	return 0
}

func EstoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	Pablo := new(hombre)
	Pablo.esHombre = true
	Pablo.vivo = true
	HumanoRespirando(Pablo)

	Maria := new(mujer)
	Maria.vivo = true
	HumanoRespirando(Maria)

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalRespirando(Dogo)
	totalCarnivoros += AnimalCarnivoro(Dogo)

	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)

	fmt.Printf("Dogo esta vivo: %t \n", EstoyVivo(Dogo))
}
