package main

import (
	"Github/Ch1nolas/Curso-Go/defer_panic"
)

func main() {
	defer_panic.EjemloPanic()
}

/* Códigos de pruebas
	numero, numero2 := ejercicios.ConviertoaNumero("2", "1000")
	fmt.Println(numero)
	fmt.Println(numero2)


	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)


if os := runtime.GOOS; os == "linux" || os == "OS X." {
	fmt.Println("Esto no es windows, es ", os)
} else {
	fmt.Println("Esto es windows, es")
}

switch os := runtime.GOOS; os {
case "linux":
	fmt.Println("Esto es Linux")
case "darwin":
	fmt.Println("Esto es Darwin")
default:
	fmt.Printf("%s \n", os)
}

	numero, texto := ejercicios.ConvNumerico("fff")
	fmt.Println(numero)
	fmt.Println(texto)

	fmt.Println(ejercicios.TabladeMultiplicar())

	Manejo de archivos
	files.LeoArchivo()
	files.SumaTabla()
	files.GrabaTabla()

	Funciones
	funciones.Exponencia(2)
	funciones.LlamarClosure()
	funciones.Exponencia()

	Arreglos y Slices
	arreglos_lices.MuestroArreglos()
	arreglos_lices.MuestroSlice()
	arreglos_lices.Capacidad()

	Mapas
	mapas.MostrarMapas()

	Modelos
	users.AltaUsuario()


	Interfaces
	e "Github/Ch1nolas/Curso-Go/ejer_interfaces" //Se puede agregar un prefijo para los imports

	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
*/
