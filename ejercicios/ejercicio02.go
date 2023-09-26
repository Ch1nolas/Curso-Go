package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* Mi código */
var numero int
var num int
var err error

func Tablademultilplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese el numero:")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error, no has ingresado un  numero")
				continue
			}
		}
		numero = num
		break
	}

	fmt.Println("Tu numero es:", numero)

	for i := 0; i <= 10; i++ {
		resultado := numero * i
		fmt.Println(resultado)
	}
}

// Código del curso
/* func Tablademultilplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {
		fmt.Println("Ingrese un numero:")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
	}
} */
