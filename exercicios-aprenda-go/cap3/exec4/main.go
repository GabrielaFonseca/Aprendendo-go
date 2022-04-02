package main

import "fmt"

type aprenda int

var x aprenda

func main() {
	fmt.Println("O valor de x é:", x)
	fmt.Printf("O tipo de x é: %T\n", x)
	x = 42
fmt.Println("O novo valor atribuido a x é:", x)
}
