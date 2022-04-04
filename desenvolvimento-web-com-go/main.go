package main

import "fmt"

func main(){
	ok, err := say("Aprendendo Go")
	if err != nil { //Se der erro no linux, use:  == e não != 
		panic(err.Error())
	}
	switch ok {
	case true: 
	fmt.Println("GOOOOOO")
	default: 
	fmt.Println("NO GOOOOOO")
	}
}

func say(what string)(bool, error){
	if what == " "{ // Se o what for uma string vazia, retorna o erro "Empty String"
		return false, fmt.Errorf("Empty String")
	}
	fmt.Println(what)
	return true, nil
}
