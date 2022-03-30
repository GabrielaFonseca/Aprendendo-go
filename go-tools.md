# Ferramentas do Go

O Go possui diversas ferramentas (_tools_) que são nativas (embutidas) que podemos utilizar durante o desenvolvimento do software e que podem ser consultadas a qualquer momento, caso esteja trabalhando sem internet (offline).

Utilize `go --help` ou `go help commando`

### Go doc

O `go doc` é um sistema de documentação onde você encontra informações sobre o pacote e as funções que deseja esta usando.

`go doc package-name` ou `go doc package-name`

O pacote `fmt` onde a `Println` é uma função.
Por exemplo: `go doc fmt.println`

O pacote `time` onde a `now` é uma função.
Por exemplo: `go doc time.now `

### Go run

\_Tradução: Run siginifica Executar.

Quando estamos desenvolvendo é possivel rodar e testar o programa com `go run`, sem precisar compilar todo o nosso codigo novamente e gerando um novo arquivo-binario, já que neste caso estamos testando o nosso codigo, sendo uma forma rapida de interação que nos permite validar erros, bugs e etc quando estavamos desenvolvendo.

O commando `go run programa.go` não vai gera um executavel e é preciso rodar o commando `go build programa.go`.

### Go build

\_Tradução: Build siginifica Construir/Criar.

Quando escrevemos um programa em GO, escrevemos de uma forma onde os desenvolvedores (_pessoas_) vão entender o código escrito, mas para que o computador entendar a instrução do nosso programa precisamos `compilar`em uma linguagem que o computador conheça como 0's e 1's(_Binarios_).

O Go tem um compilador é um software que converte o código em Go, em uma forma que o computador entende.

Esse código traduzido, criando um arquivo executável ou arquivo/binário.

Podendo então ser executável na maquina/servidor, onde o programa escrito realizar as instruções presentes no codigo.

Para traduzir o codigo que escrevemos para o linguagem de maquina, usamos o `build`.

Para compilar um programa em Go `go build main.go`

```
go build main.go
ls -la
main.go
main
./main.go
Aprendendo GO
```

Onde irá gerar um binario compilado `main.go`
