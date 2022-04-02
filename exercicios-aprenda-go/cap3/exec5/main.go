package main

import "fmt"

type aprenda int

var x aprenda
var y int

func main() {
	fmt.Println("O valor de x é:", x)
	fmt.Printf("O tipo de x é: %T\n", x)
	x = 42
	fmt.Println("O novo valor atribuido a x é:", x)
	y = int(x) 
	fmt.Println("O valor de y é:", y)
	fmt.Printf("O tipo de y é: %T\n", y)
}
