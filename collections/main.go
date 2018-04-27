package main

import (
	"fmt"
)

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}
	for _, v := range numeros {
		fmt.Println(v)
	}

	nombresSlice := []string{"nicolas", "francisca", "camilo", "jorge", "sylvia"}
	for _, v := range nombresSlice {
		fmt.Println(v)
	}

	nuevosNombes := []string{"Alonso", "Antonio", "Roberto"}
	copyNames := make([]string, 5)
	for k, v := range nombresSlice {
		copyNames[k] = v
	}
	fmt.Println(copyNames)
	fmt.Println(nuevosNombes)

	d := append(copyNames, nuevosNombes[0], nuevosNombes[1], nuevosNombes[2])
	fmt.Println(d)
	fmt.Printf("tamaño del slice: %d", len(d))

	// Mapas en go
	dict := make(map[string]string)
	dict["go"] = "golang"
	dict["cs"] = "CSharp"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "Javascript"

	for k, v := range dict {
		fmt.Printf("Llave: %s. Valor: %s\n", k, v)
	}
}
