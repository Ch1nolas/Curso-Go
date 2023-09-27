package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje fnal")

	fmt.Println("Este es el segundo mensaje")
}

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
