package main

import "fmt"

func main() {
	saludar("Jhonny", 29)
}

func saludar (nombre string, edad uint8){
	fmt.Printf("hola %s, ya tienes %d años!\n", nombre, edad)
}