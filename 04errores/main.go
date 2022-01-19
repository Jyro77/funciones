package main

import (
	"errors"
	"fmt"
)

func main() {
	//valor nulo de los errores es nil
	// variables de errores se declaran como err "habitualmente"
	// errors es un paquete a importar

	// err := errors.New("Un mensaje de error")
	r, err := div(100, 0)
	// fmt.Printf("%T\n", err) tipo de error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
		fmt.Println("Resultado es:", r)

}

func div(x, y float64) (resultado float64, err error) {
	if y == 0 {
		err = errors.New("No se puede dividir entre cero")
	} else {
		resultado = x / y
	}
	return 
}