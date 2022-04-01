package main

import (
	"fmt"
)
var x int
var y string
var z bool
//Retornando o valor zero, quando não atribuido (inicializado) um valor a uma variavel.
func main (){
	fmt.Printf("x: é  %v, --> (%T)\n", x, x)
	fmt.Printf("y: é  %v, --> (%T)\n", y, y)
	fmt.Printf("z: é  %v, --> (%T)\n", z, z)
}
