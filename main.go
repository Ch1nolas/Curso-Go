package main

import (
	"Github/Ch1nolas/Curso-Go/funciones"
)

func main() {
	funciones.LlamarClosure()
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

	files.LeoArchivo()
	files.SumaTabla()
	files.GrabaTabla()

*/
