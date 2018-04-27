package main

import (
	"fmt"

	"github.com/wainola/go_web/packages/strcon"
)

func main() {
	s := strcon.SwapCase("Ghopher")
	fmt.Println("Converted string is: ", s)
}
