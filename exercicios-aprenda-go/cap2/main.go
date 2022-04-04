/* package main

import "fmt"

var w int
var x float64
var y string
var z bool


func main (){
	fmt.Printf("x: é  %v, --> (%T)\n", w, w)
	fmt.Printf("x: é  %v, --> (%T)\n", x, x)
	fmt.Printf("y: é  %v, --> (%T)\n", y, y)
	fmt.Printf("z: é  %v, --> (%T)\n", z, z)
} */
package main

import "fmt"

type comida int
var x comida = 10

func main (){

	fmt.Printf("x: é  %v, --> (%T)\n", x)

}
