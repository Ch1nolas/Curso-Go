package ejer_interfaces

import (
	"Github/Ch1nolas/Curso-Go/interfaces"
	"fmt"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
