# Tipos

Os tipos em Go são estaticos, quando declarado um determinado tipo é para sempre, não mudando estado, sendo imutavel.

> Toda variavel delcarada em Go precisa ser utilizada, assim como os seus pacotes.

`var` Se trata de um valor ou tipo declarado aa variavel, Ex: `var x := 2` ou `var x int`.

```
package main

import "fmt"

var x int //Somente deve ser declarada a nivel do package, caso não atribuido o valor inicial e somente o tipo.

   func main (){
		x = 10 //Somente no codeblock é possivel declarar um valor.
		fmt.Println(x)
}
```

- Somente deve ser delcarada no package-level-scope, quando declarada no bloco de instrução não interagem com nenhuma outra função.

- O valor atribuido a `var` quando não declarado via pacotes, só pode ser declarado dentro do bloco de instruções.

### Tipos de dados Primitivos

Podemos ter variaveis do tipo `numero`, `texto` e `boleano`.

`int` Se trata de um valor de numero inteiro, Ex: `1`.
`float` Se trata de um valor flutuante contendo `,` ou ponto `.`, Ex: `Pi 3.14` ou `R$0,50`.
`boolean`Se trata de uma condicão sendo verdadeira ou falsa, Ex: `true` ou `false`.

```
package main

import "fmt"

func main (){
	x := 0 //Numero
	y := "string" //texto
	z := true  //boolean
	fmt.Println(x, y, z)
}
```

- É necessario declara o tipo da variavel

Variaveis = delcarar uma variavel a ser usada.
Instruções = são ações que o programa vai realizar.

### Tipos de dados Compostos

São os tipos compostos criados pelo usuario que compoe o uso dos tipos primitivos.

`slice`
`array`
`struct`
`map`

### Valor Zero

Valor zero é quando uma variavel e declarada e não inicializada tendo por padrão o valor Zero (0).

Declaração - Separar o endereço a ser utilizado na memoria pela variavel, Ex: `var x int`.
Inicialização - Atribui um Valor a variavel, Ex: ` x = 10`.
Atribuição - Dar um novo valor a determinada variavel, Ex: ` x = 20`.

- É a declaração de uma variavel sem atribuir um valor.

```
package main

import "fmt"

var w int
var x float64
var y string
var z bool


func main (){
	fmt.Printf("w: é  %v, --> (%T)\n", w, w)
	fmt.Printf("x: é  %v, --> (%T)\n", x, x)
	fmt.Printf("y: é  %v, --> (%T)\n", y, y)
	fmt.Printf("z: é  %v, --> (%T)\n", z, z)
}

```

Quais são os tipos de declaração que retorna valor zero?

> int: 0
> float: 0.0
> boolean: false
> string: ""
> pointer, functions, interfaces, slices, channels, maps: nil

## Operadores

### Operador curto de declaração

\_Traducão: Gopher significa marmota.

> O operador é conhecido como Gopher.

O operador atribui um tipo da variavel com base no seu valor declarada no codigo, se tratando da tipagem automatica.

`=` Usado para atribuir valores a variaveis já existentes.
`:=` Usado para declarar valores a novas variaveis ou pelo menos ter uma nova variavel a ser declarada quando usada com uma variavel que já possui um valor.

```
package main

import "fmt"

func main (){
	x := 1 //Numero
	y := "string" //letras
	x = 2
	x, z := 3, 10 //O valor de X mudar, pois depende da declaração de uma nova variavel --> Z
	fmt.Printf("x: é  %v, --> (%T)\n", x, x)
	fmt.Printf("y: é  %v, --> (%T)\n", y, y)
	fmt.Printf("z: é  %v, --> (%T)\n", z, z)
}
```

- É utilizado para declarar uma variavel.
- Marmota somente funciona dentro do bloco da instrução (_codeblock_)

`%v`
`%T`

### Operadores Logicos

`&&` (AND)
`||` Comparação (OR), compara o valor
`!` Negação (NOT), inverte o valor

### Operadores aritméticos

`+`
`-`
`*`
`/`
`%`
`++`
`--`

### Operadores relacionais

`==` É igual o valor, Ex `10 == 100 = Falso`.
`!=` É diferente do valor, , Ex `10 != 100 = Verdadeiro`.
`>` Maior o valor, Ex `10 > 100 = Falso`.
`<` Menor o valor, Ex `10 < 100 = Verdadeiro`.
`>=` Maior ou igual o valor, Ex `10 >= 100 = Falso`.
`<=` Menor ou igual o valor, , Ex `10 <= 100 = Verdadeiro`.

### Tatamento de Erros

response, err
result, err

`_` = descartar as informaões recebidas.

### Fazendo comentarios no codigo em Go

Ao comentar determinada linha ou bloco do código, o compilador vai ignorar determinada instrução.

- Comentários no inicio da linha `//Linha` sendo ignorado pelo compilador:

```
package main

import (
	"fmt"
)

func main() {
    //fmt.Println("Compile o Programa)
	/*fmt.Println("Go")
    fmt.Println("Go Go Go")*/
    fmt.Println("GOOOOOOOLAAAANG")//Go
}
```

- Comentários depois de uma instrução presente na linha `instrução //Linha` sendo ignorado pelo compilador:

```
package main

import (
	"fmt"
)

func main() {
    fmt.Println("GOOOOOOOLAAAANG")//Go
}

```

- Comentários em bloco que podem abranger várias linhas, começando com `/*` e terminando com `*/` e pode envolve tudo dentro do bloco (incluindo novas linhas):

```
package main

import (
	"fmt"
)

func main() {
    //fmt.Println("Compile o Programa)
	/*fmt.Println("Go")
    fmt.Println("Go Go Go")*/
    fmt.Println("GOOOOOOOLAAAANG")//Go
}

```

---

# Criando tipos

Ë possivel customizar um tipo, criando
type

Conversão de tipos
