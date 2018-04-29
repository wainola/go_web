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

// Metodo para persona
func (p Persona) PrintName() {
	fmt.Printf("\n%s %s\n", p.Nombre, p.Apellido)
}

func (p Persona) PrintDetails() {
	fmt.Printf("[Added: %s, Email: %s, Nombre: %s %s]\n", p.Added, p.Email, p.Nombre, p.Apellido)
}

// set metod => se utilizar un puntero
func (p *Persona) CambioNombre(nuevoNombre string) {
	p.Nombre = nuevoNombre
}

type Admin struct {
	Persona
	Roles []string
}

func main() {
	// Instanciado tipos
	var p Persona
	p.Nombre = "Nicolas"
	p.Apellido = "Riquelme"
	p.Email = "nrriquelme@gmail.com"
	p.Edad = 30
	p.Added = time.Date(2018, 04, 27, 0, 0, 0, 0, time.UTC)

	var adm = Admin{
		Persona: p,
		Roles:   []string{"jefe", "administrador de sistemas", "programador"},
	}

	fmt.Println(p)
	fmt.Println(adm)

	// Coleccion de tipos.
	curso := make([]Persona, 3)
	curso[0] = Persona{
		Nombre:   "Camilo",
		Apellido: "Riquelme",
		Email:    "camilo@mail.com",
		Edad:     28,
		Added:    time.Date(2018, 04, 27, 0, 0, 0, 0, time.UTC),
	}
	curso[1] = Persona{
		Nombre:   "Sylvia",
		Apellido: "Guzman",
		Email:    "olmue2@gmail.com",
		Edad:     56,
		Added:    time.Date(2018, 04, 27, 0, 0, 0, 0, time.UTC),
	}

	curso[2] = Persona{
		Nombre:   "Jorge",
		Apellido: "Aliaga",
		Email:    "jorge@gmail.com",
		Edad:     55,
		Added:    time.Date(2018, 04, 27, 0, 0, 0, 0, time.UTC),
	}

	for _, v := range curso {
		if v.Nombre == "Nicolas" {
			fmt.Println("Hay un Nicolas en el slice!")
		} else {
			fmt.Println("No hay Nicolas en el slice")
		}
	}

	curso[0].PrintDetails()
	curso[0].CambioNombre("Camilo Sebastián")
	curso[0].PrintDetails()
}
