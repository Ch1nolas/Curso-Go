package ejercicios

import (
	"strconv"
)

func ConvNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}

// Mi codigo
/* func ConviertoaNumero(texto string, texto2 string) (int, bool) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, false
	}

	numero2, err := strconv.Atoi(texto2)
	if err != nil {
		return 0, false
	}
	if numero > 100 {
		fmt.Println("El numero 1 es mayor a 100")
	} else {
		fmt.Println("El numero 1 es menor a 100")
	}
	if numero2 > 100 {
		fmt.Println("El numero 2 es mayor a 100")
	} else {
		fmt.Println("El numero 2 es menor a 100")
	}

	return numero, true

} */
