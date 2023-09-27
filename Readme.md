# Mi Documentación

## Iniciar un Proyecto Go

Para iniciar un proyecto en Go, utiliza el siguiente comando:

```bash
go mod init Github/Ch1nolas/Curso-Go (path)
```

## Estructura de un Archivo Go

Es fundamental que la primera línea de un archivo Go indique a qué paquete pertenece. Puedes hacerlo de la siguiente manera:

```go
package main // Es cuando es el archivo main y no está en ninguna subcarpeta
package variables // En este caso es un archivo que está dentro de la subcarpeta variables
```

## Importar paquetes

Para importar paquetes en Go, utiliza el siguiente formato:

```go
import (
"fmt" // Es un paquete que sirve para mandar mensajes a la consola
)
```

## Funciones

Aquí tienes un ejemplo de una función básica en Go:

```go
func main() {
fmt.Println("Hola Mundo")
}
```

## Ejecutar código

Para ejecutar un programa Go, utiliza el comando:

```bash
go run main.go
```

## Compilar código

Para compilar un programa Go, utiliza el comando:

```shell
go build main.go
```

## Variables públicas

Las variables públicas en Go deben comenzar con una letra mayúscula. Por ejemplo:

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

Ejemplo básico de un condicional 'if' en Go:

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

Operadores condicionales en Go:

- `==` (igual a)
- `!=` (no igual a)
- `<` (menor que)
- `>` (mayor que)
- `<=` (menor o igual que)
- `>=` (mayor o igual que)
- `&&` (y lógico)
- `||` (o lógico)

Método 'switch' en Go:
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

## Manejo de archivos

**Revisar y hacer teoría Clase 17**

## Funciones Anónimas

Las funciones anónimas en Go se crean sin un nombre y se pueden asignar a variables. Aquí tienes un ejemplo:

```go
func Calculos() {
  suma := func(numero1 int, numero2 int) int {
    return numero1 + numero2
  }
  fmt.Println(suma(10, 25))
}
```

## Closures

Los closures en Go son funciones que capturan variables del entorno en el que se crean. Aquí tienes un ejemplo:

```go
func tabla(valor int) func() int {
  numero := valor
  secuencia := 0
  return func() int {
    secuencia++
    return numero * secuencia
  }
}

func LlamarClosure() {
  tabladel := 2
  MiTabla := tabla(tabladel)
  for i := 1; i <= 10; i++ {
    fmt.Println(MiTabla())
  }
}
```

## Recursion

La recursión en Go se refiere a la capacidad de una función de llamarse a sí misma. Aquí tienes un ejemplo:

```go
func Exponencia(valor int) {
  if valor > 100000000 {
    return
  }
  fmt.Println(valor)
  Exponencia(valor * 2)
}
```

## Arreglos

Los arreglos en Go son colecciones de elementos de un solo tipo de datos. Aquí tienes un ejemplo de declaración y uso de arreglos en Go:

```go
var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {
  tabla[7] = 33
  tabla[2] = 54
 
  tabla2 := [10]string{"Chinolas", "Matias"}
  fmt.Println(tabla)
  fmt.Println(tabla2)

  for i := 0; i < len(tabla); i++ {
    fmt.Println(tabla[i])
  }
 
  matriz[7][24] = 15
  fmt.Println(matriz)
}
```

## Slices

Los slices en Go son estructuras de datos dinámicas que permiten una capacidad flexible. Aquí tienes un ejemplo de declaración y uso de slices en Go:

```go
var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
  fmt.Println(tablaS)
  porcion := arreglo[3:]   //Slice  creado desde un vector, desde la posición 3
  porcion2 := arreglo[:5]  //Slice creado desde un vector, desde la posición 0 hasta la 5
  porcion3 := arreglo[6:7] //Slice creado desde un vector, desde la posición 6 hasta la 7
  fmt.Println(porcion)
  fmt.Println(porcion2)
  fmt.Println(porcion3)
}
```

Lo potente de los slices es que puedes ajustar su capacidad dinámicamente utilizando la función 'make'. Aquí tienes un ejemplo:

```go
func Capacidad() {
  elementos := make([]int, 5, 20)
  fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
 
  nums := make([]int, 0, 0)
  for i := 0; i < 100; i++ {
    nums = append(nums, i)
  }
  fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
```

## Para los textos

- `%d` Para armar un texto con una variable entera
- `%s` Para armar un texto con una variable de tipo string
- `%t` Para armar un texto con una variable de tipo bool

## Mapas

```go
func MostrarMapas() {
  paises := make(map[string]string) // Sintaxis del mapa
  //fmt.Println(paises)

  // Poner los mapas manualmente
  paises["Mexico"] = "D.F"
  paises["Argentina"] = "Buenos Aires"
  fmt.Println(paises)
  fmt.Println(paises["Argentina"])


  //Sintaxis del mapa más rápida y con menos código
  campeonato := map[string]int{
    "Barcelona":    39,
    "Real madrid":  38,
    "Chivas":       37,
    "Boca Juniors": 30,
  }
  fmt.Println(campeonato)


  //Es como el foreach, para recorrer el mapa
  for equipo, puntaje := range campeonato {
      fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
    }
   
  // Para borrar un elemento del mapa
  delete(campeonato, "Real madrid")
  fmt.Println(campeonato)
 
  puntaje, existe := campeonato["Chivas"]
  fmt.Printf("El puntaje captura es %d, y el equipo existe = %t \n", puntaje, existe)
}
```

## Estructura

### Modelos

```go
// Modelo para los usuarios
type User struct {
  Id        int
  Name      string
  CreatedAt time.Time
  Status    bool
}

//Definimos el modelo para luego crearlo desde el comando
func (usuario *User) AddUser(id int, name string, createdAdd time.Time, status bool) {
  usuario.Id = id
  usuario.Name = name
  usuario.CreatedAt = createdAdd
  usuario.Status = status
}
```

Función para crear el usuario

```go
import "Github/Ch1nolas/Curso-Go/modelos" //Hay que importar los modelos que tengamos

func AltaUsuario() {
  u := new(modelos.User)
  u.AddUser(10, "Chinolas", time.Now(), true)
  fmt.Println(u)
}
```

## Interfaces

**Revisar la clase y completar la documentación**

```go
//Sintaxis de las interfaces
type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
	EstaVivo() bool
}
```

```go
// Estructura del modelo y autocompletado de la interfaz
type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	comiendo   bool
	pensando   bool
	vivo       bool
}

func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Comer()       { h.comiendo = true }
func (h *Hombre) Pensar()      { h.pensando = true }
func (h *Hombre) Sexo() string { return "Hombre" }
```

```go
// Como lo fundamentamos en el main
    Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
```

## Defer y Panic

**Defer**

```go
func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje fnal")
	fmt.Println("Este es el segundo mensaje")
}
```

**Panic**
Sirve para abortar el código y dejar un mensaje

```go
func EjemloPanic() {
	a := 1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}
```

**Recover**

```go
func EjemloPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó un panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}
```
