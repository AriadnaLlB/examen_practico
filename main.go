package main

import "fmt"

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func main() {
	resultadoSuma := suma(2, 2)
	fmt.Println("El resultado de la suma es:", resultadoSuma)

	resultadoResta := resta(4, 5)
	fmt.Println("El resultado de la resta es:", resultadoResta)
}
