# Estrutura de Codigo em Go

Entendendo a estrutura do codígo escrito em Go.

## Pacotes

\_Tradução: Packages siginifica Pacotes.

Todo programa escrito em Go é composto por `Packages`, onde um programa começa a ser executado pelo pacote principal `main`, a ser declarado na primeira linha do codigo, **`package main`**.

```
 package main   <---
 import "fmt"
 func main()
 {
  fmt.Println("Aprendendo Go!")
 }
```

---

## Imports

_Tradução: Imports siginifica Importações._

Assim como Pacote main, o Go possui outros pacotes que podem ser importados e utilizados no código para realizar diferentes tarefas.

Um dos pacotes mais populares é o `fmt`, que significa formato e tem a funcionalidade de entrada e saída `Input/Output` de informações/dados.

É possivel importar diversos pacotes ao mesmo tempo, usando parenteses:

```
import (
  "fmt",
  "time",
)
```

É possivel declarar o uso de um ou mais pacotes no `import`

- Usando um unico `import`:

```
package main

import ( "fmt"
         "erros"
)

func main() {
	fmt.Println("Aprendendo Go")
}
```

- Usando multiplos `import`:

```
package main

import "fmt"
import "errors"

func main() {
	fmt.Println("Hello World")
}
```

Quando o codigo não está fazendo o uso de um determinado Pacote/Lib na função, o Go remove a lib do import.

#### Import e Alias

Podemos usar um `alias` para o pacote a ser utlizado, somente incluindo um nome antes do biblioteca.

```
package main

import (
	"fmt"
	t "time"
)

func main() {
	fmt.Println(t.Now())
}
```

Sendo possivel chamar o alias através de uma função a ser usada **`fmt.Println(t.Now())`** ao invés de `fmt.Println(time.Now())`.

```
package main

import (
	"fmt"
	t "time"
)

func main() {
	fmt.Println(t.Now())
}
```

---

## Função

\_Tradução: Function siginifica Função.

Uma `func` é um bloco de instruções de codigo que podem ser usadas/reutilizadas ao longo das tarefas.

```
 package main   <---
 import "fmt"
 func main()
 {
  fmt.Println("Aprendendo Go!")
 }
```

- `func` é o inicio da função a ser declarada.
- `func name` é declarado o nome da função.
- Após o nome da função vem um par de parênteses `()` e um conjunto de chaves `{}`.
- Dentro do conjunto de chaves `{ fmt.Println("Aprendendo Go!") }` possui a instrução do que o codigo íra fazer.

Existem dois tipos de funções utlizadas no pacotes `Publica` e `Privada`.

- `Funções Publicas` são escritas com letra maiscula `fmt.Println`.
- `Funções Privadas`,são escritas com letras minusculas.

As funções podem receber multiplos valores

```
package main

import (
	"fmt"
)

func swap(x, y string) (string, string){ //Função de inverter a frase
	return y, x
}

func main() {
	a, b := swap("Aprendendo", "Go")
	fmt.Println(a, b)
}
```

---

### Bibliotecas

\_Tradução: Library siginifica Biblioteca.

---

## Erros

É comum a tratativa de erros durante o desenvolvimento do codigo em Go.

## GoRoutines

São rotinas que trabalham em paralelo/ocorrencia durante a execução de uma determinada tarefa/função.

`Go`

## Canais

São espaços compartilhados em memoria

## Cross Compilation

Por ser uma linguagem focada em sistemas distribuidos Go suporta diversas versões de SO (_ Versões de Sistema Operacional) & CPU (\_Arquitetura de 32/64 bits_) sendo possivel criar executáveis para diferentes plataforma de acordo com a arquitetura escolhida, sendo diferente da versão que o compilador utilizando no momento e que executa o codigo e/ou cria um novo arquivo-binario.

Em [Ferramentas do Go](go-tools.md) você irá encontrar uma seção sobre o assunto.
