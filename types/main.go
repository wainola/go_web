package main

import (
	"fmt"
	"time"
)

type Persona struct {
	Nombre   string
	Apellido string
	Email    string
	Edad     int64
	Added    time.Time
}

func main() {
	// Instanciado tipos
	var p Persona
	p.Nombre = "Nicolas"
	p.Apellido = "Riquelme"
	p.Email = "nrriquelme@gmail.com"
	p.Edad = 30
	p.Added = time.Date(2018, 04, 27, 0, 0, 0, 0, time.UTC)

	fmt.Println(p)

	// Coleccion de tipos.
	curso := []Persona{}
	curso[0] = Persona{
		Nombre: "Camilo",
		Apellido: "Riquelme",
		Email: "camilo@mail.com",
		Edad: 28,
		Added: time.Date(2018,04,27,0,0,0,0,time.UTC)
	}
	curso[1] = 
}
