package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		estado, texto := variables.ConviertoaTexto(1588)
		fmt.Println(estado)
		fmt.Println(texto)
	*/

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

}
