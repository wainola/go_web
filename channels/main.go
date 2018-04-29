package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// declarando canal
	count := make(chan int)

	// añadendo 2 working groups por las dos goroutinas que ejecutaremos
	wg.Add(2)

	fmt.Println("Comenzando goroutinas")
	// lanzando goroutina con la etiqueta A
	go printCounts("A", count)
	// lanzando gorotuna con la etiqueta B
	go printCounts("B", count)
	fmt.Println("Canal comienza")
	// ingresando datos al canal
	count <- 1
	// Esperando que goroutinas terminen
	fmt.Println("Esperando que terminen")
	wg.Wait()
	fmt.Println("\nTerminando el programa")
}

func printCounts(label string, count chan int) {
	// programando para indicarle al wg que ya hemos terminado
	defer wg.Done()
	for {
		// Recibiendo mensaje desde el canal
		val, ok := <-count
		if !ok {
			fmt.Println("El canal fue cerrado")
			return
		}
		fmt.Printf("Contador: %d recibido desde %s \n", val, label)
		if val == 10 {
			fmt.Printf("Canal cerrado desde %s \n", label)
			// cerrando el canal
			close(count)
			return
		}
		val++
		// enviando contador hacia la otra goroutina
		count <- val
	}
}
