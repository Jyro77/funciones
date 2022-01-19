package main

import "fmt"

func main() {
	// funcion anonima que se ejecuta por si misma
	/* 	func (){

	   	}() */
	// Tambien puede ser guardada en una variable y ejecutarla mas adelante
	/* 	anonima := func() {

	}
	   	anonima() */
	resultado1 := secuencial()
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())
	fmt.Println(resultado1())

	fmt.Println("-----------------------------------------")
	
	resultado2 := secuencial()
	fmt.Println(resultado2())
}

func secuencial() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}