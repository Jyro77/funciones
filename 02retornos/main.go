package main

import "fmt"

func main() {
	var n1, n2 int8
	n1 = 15
	n2 = 3
	r := sum(n1, n2)
	fmt.Println(r)

	edad := uint8(3)
	e := edadType(edad)
	fmt.Println(e)

	n := []int8{52, 63, 47, 5, 3, 7, 6, 100, 2, -5, 60, 82, -52, 120}
	fmt.Println(minymax(n))
}

// la firma de la funcion
func sum(x , y int8) int8 {
	return x + y
}

func edadType(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "Niñez"
	case edad < 18:
		tipo = "Adolecencia"
	default:
		tipo = "Adulto"
	}
	return tipo
}

// Para devolver mas de un valor se tiene que hacer un segundo parentesis y declarar el tipo de los valores que estará retornando
func minymax(num []int8) (min int8, max int8) {
	for _, v := range num {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}