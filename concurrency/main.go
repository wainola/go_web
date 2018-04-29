package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg es usado para esperar a que el programa finalice las rutinas
var wg sync.WaitGroup

func main() {
	// añadiendo 2, puesto que tendremos dos goroutines operando.
	wg.Add(2)

	fmt.Println("Iniciando goroutines")
	// lanzando la goroutine con la etiqueta A
	go printCounts("A")
	// lanzando la goroutine con la etiqueta B
	go printCounts("B")
	// esperar que las goroutines terminen
	fmt.Println("Esperando finalizacion")
	wg.Wait()
	fmt.Println("\n Terminando el programa")

	wg.Add(2)
	fmt.Println("Lanzando dos nuevas goroutines")
	go processDb()
	go sayNombre("Camilo Riquelme")
	fmt.Println("Esperando finalizacion de goroutines")
	wg.Wait()
	fmt.Println("\nTerminando el programa")

}

func printCounts(label string) {
	// indicando al grupo de espera que la funcion finalizo
	defer wg.Done()
	// Esperando aleatoriamente
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Cuenta: %d from %s\n", count, label)
	}
}

func processDb() {
	defer wg.Done()
	fmt.Println("Procesando database")
	for c := 1; c <= 20; c++ {
		fmt.Printf("Procesando database %d\n", c)
	}
}

func sayNombre(name string) {
	defer wg.Done()
	fmt.Printf("Tu nombre es %s\n", name)
}
