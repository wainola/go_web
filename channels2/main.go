package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	precioArriendo := make(chan int)

	wg.Add(2)
	fmt.Println("Comenzando el programa con las goroutinas")
	go calcularPrecio("A", precioArriendo)
	go calcularPrecio("B", precioArriendo)
	fmt.Println("Iniciando el canal con el valor 1000")
	precioArriendo <- 1000
	fmt.Println("Esperando que se cierre el canal para terminar el programa")
	wg.Wait()
	fmt.Println("\nTerminando el programa")
}

func calcularPrecio(etiqueta string, c chan int) {
	// avisar al wg cuando hemos terminado
	defer wg.Done()

	for {
		val, ok := <-c
		if !ok {
			fmt.Println("El canal fue cerrado. Ya se hizo todo el procesamiento")
			return
		}
		fmt.Println("Calculando precio del arriendo")
		arriendo := 5000 + val
		fmt.Printf("El valor del arriendo es %d \n", arriendo)
		if val == 3000 {
			fmt.Printf("Cerrando el canal desde %s\n", etiqueta)
			close(c)
			return
		}
		val += 1000
		// insertando en el canal
		c <- val
	}
}
