package main

import (
	"fmt"
	"time"
)

type Usuarios interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName, LastName string
	Added               time.Time
	Email, Location     string
}

func (p Person) PrintName() {
	fmt.Printf("\n nombre de la persona: %s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[nombre: %s %s, Email: %s, Locacion: %s, Añadido: %s\n]", p.FirstName, p.LastName, p.Email, p.Location, p.Added.String())
}

func main() {
	nicolas := Person{
		"Nicolas",
		"Riquelme",
		time.Date(1987, time.December, 29, 14, 0, 0, 0, time.UTC),
		"nrriquelme@gmail.com",
		"Santiago",
	}

	camilo := Person{
		"Camilo",
		"Riquelme",
		time.Date(1990, time.July, 27, 12, 0, 0, 0, time.UTC),
		"camilo@mail.com",
		"Santiago",
	}

	usuarios := []Usuarios{nicolas, camilo}
	for _, v := range usuarios {
		v.PrintName()
		v.PrintDetails()
	}
}
