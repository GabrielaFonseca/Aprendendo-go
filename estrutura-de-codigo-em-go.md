# Estrutura de Codigo em Go

Entender a estrutura do codígo em Go

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

---

## Pacotes

\_Tradução: Packages siginifica Pacotes.

Todo programa escrito em Go é composto por `Packages`, onde um programa começa a ser executado no pacote principal `main`, a ser declarado na primeira linha do codigo, **`package main`**.

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

#### Boas praticas para o dia a dia

- Utilize o multiplo import quando usar mais de um package/lib.
- Utilize o álias com moderação, pense no quanto está legivel (facil de ler) para outra pessoa quando for trabalhar em time.

### Bibliotecas

\_Tradução: Library siginifica Biblioteca.

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
