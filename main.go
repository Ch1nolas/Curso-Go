package main

import (
	"Github/Ch1nolas/Curso-Go/variables"
	"fmt"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
