# Mi documentación

Para iniciar un proyecto en go

```bash
go mod init Github/Ch1nolas/Curso-Go (path)
```

A la hora de crear un archivo GO es fundamental que la primera línea sea el package "nombre de la carpeta a donde corresponde"

```go
package main // Es cuando es el archivo main y no está en ninguna subcarpeta
package variables // En este caso es un archivo que está dentro de la subcarpeta variables
```

## Importar paquetes

```go
import (
"fmt" // Es un paquete que sirve para mandar mensajes a la consola
)
```

## Funciones

```go
func main() {
fmt.Println("Hola Mundo")
}
```

## Ejecutar código

```bash
go run main.go
```

## Compilar código

```shell
go build main.go
```

## Variables públicas

La variable tiene que comenzar con **mayúscula**
Declaración de variables

```go
var intcomun int
```

Para importar las carpetas hay que poner el path completo, en este caso es así:

```go
import (
"Github/Ch1nolas/Curso-Golang/variables"
)
```

## Condicionales

```go
//Ejemplobásico del if
os := runtime.GOOS
if os=="Linux."{
fmt.Printlm("Esto es un linux")
} else {
fmt.Println("Esto es un windows")
}

//If pero en menos líneas, se separan con un ;
if os := runtime.GOOS; os=="Linux."{
fmt.Println("Esto es un linux")
} else {
fmt.Println("Esto es un windows")
}
```

Operadores condicioneales

```go
==
->
-<
=>
=<
&&
||
```

Método Switch
Te permite dejar por defecto ciertas consignas para que si se cumple una, se realize su respectiva reacción. Y el default sirve para todas las otras opciones que no hemos puesto, en este ejemplo sirve para los otros tipos de sistemas operativos que no hemos mencionado en los case. **Se recomienda utilizar para cuando se tiene más de 2 condiciones en el if**

```go
switch os := runtime.GOOS; os {

  case "linux":

    fmt.Println("Esto es Linux")

  case "darwin":

    fmt.Println("Esto es Darwin")

  default:

    fmt.Printf("%s \n", os)

  }
```

En GO se puede hacer que cuando se te pidan dos valores uno tengo un valor nulo

```go
num, __ := strconv.Atoi(texto)
```

El paquete bufio sirve para la lectura del teclado, **Para tener en cuenta a la hora del Proyecto Integrador**

# Iteraciones

## Método for

```go
for i :=0; i<10; i++ {
	fmt.Println(i)
}
```

Para abortar se utiliza el **break** (corta todo el programa)

```go
func Iterar() {
  for i := 0; i < 10; i++ {
    if i == 6 {
      break
    }
    fmt.Println(i)
  }
}
```

```bash
1
2
3
4
5
Se corta el programa
```

Y para continuar se utiliza el **continue**

```go
func Iterar() {
  for i := 0; i < 10; i++ {
    if i == 6 {
      continue
    }
    fmt.Println(i)
  }
}
```

```shell
0
1
2
3
4
5
7
8
9
```

Donde acá omite al número 6 y sigue con el for
