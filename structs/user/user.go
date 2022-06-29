package user

import "time"

// Definición de estructura - equivalente a definir una clase
// Como se exporta, debe tener el nombre con la primera letra en mayúscula
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaAlta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}
