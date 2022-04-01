package main

import (
	"fmt"
)
var x = 42
var y = "James Bond"
var z = true

func main (){
	s := fmt.Sprintf(" x é: %v\n y é: %v\n z é: %v\n ", x, y, z)
	fmt.Println(s)

}
