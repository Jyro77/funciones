package main

import "fmt"

func main() {
	saludar("Jose", "Marcos", "Luis", "Mateo", "Pedro", "Lucas", "Juan")
}

// recibe uno o varios parametros
// solo recibe un unico parametro variatico y este debe de estar declarado al final
// Ejemplo **func saludar(edad uint8, nombre ...string)**
func saludar(nombre ...string) {
	for _, v := range nombre {
		fmt.Println("hola ", v)
	}
}