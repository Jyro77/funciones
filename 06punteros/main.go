package main

import "fmt"

// puntero es apuntar a la "Direccion de la memoria de "
func main() {
	//para obtener la direccion de la memoria de una variable se le antepone el "&" ante la variable. Ejemplo:
	a := 3
	fmt.Println(&a)
	// el valor cero de un puntero es "nil"

	fmt.Println("antes de duplicar", a)
	duplicar(&a)
	fmt.Println("despues de duplicar", a)
}

// para acceder a el valor de memoria por parametros se antepone el *... de la siguiente manera:
func duplicar(x *int) {
	//*x = *x * 2
	*x *= 2
	fmt.Println("Dentro de Duplicar", x)
}
// De esta manera trabajamos con lo que originalmente esta guardado en ese espacio de memoria
// a esto se le llama: "paso por referencias o por punteros"