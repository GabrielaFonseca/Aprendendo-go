# Entendo os pacotes e as funçõe

`fmt.Println(...interface{}) `(n int, err error`

Isso significa que `[...interface]` pode recebe multiplos números de variável de diferentes argumentos, podendo ser de qualquer tipo.

Entregando o retorno `(n int, err error` de um `numero inteiro` e o tipo do `erro`, informando o que aconteceu.

Existem dois tripos de Strings em Go que são interpretadas de formas diferentes.

`Literals` representa um valor fixo no codigo.

```
package main

import "fmt"

func main (){
	fmt.Println(`"Aqui temos um texto\
	escrito da\
	forma que\
	quisermos no codigo`)

	fmt.Print("Aqui é um texto indexado")
}
```

`Interpreted String Literals` onde o conteudo é intepretado por aspas duplas`"Uma texto contendo aspas duplas"` mantendo a indexação do texto no codigo.

`Raw String Literals` onde o conteudo é intepretad por acentos graves `"Um texto contendo \ acento grave "` mantendo a forma do texto escrito no codigo.

Todo o caracter em uma `string` é chamado de `rume literals`

#Pacotes

`P` Print
`S` String
`F` File
`ln` adiciona uma nova linha
`tf`

## Fmt

Print = não adiciona uma linah no final
Println = adiciona uma nova linha no final
PrintFn

Sprint

```
package main

import "fmt"

func main (){
	x := "Go"
	y := "lang"
	z := fmt.Sprint(x, y) // O valor é salvo em uma string

	fmt.Println(z) //Adiciona uma linha em branco

}
```
