package main

import "fmt"

func main() {

	var numero int
	var opcion string

	menu := "Ingresa el número que quieres transformar "

	fmt.Print(menu)
	fmt.Scanln(&numero)

	menu2 := "Quieres transformar el número en 1-. Factoria 2-. Fibonacci"
	fmt.Println(menu2)
	fmt.Scanln(&opcion)

	if opcion == "1" {
		optionFacto(numero)
		//Terminado
	} else if opcion == "2" {
		optionFibona(numero)
		//Terminado
	} else {
		fmt.Println("Información incorrecta. Escribe 1 o 2")

	}

}

func optionFacto(numero int) {
	fmt.Println("El numero en factorial de ", numero, " Es:")
	fmt.Println(factorial(numero))
}

func optionFibona(number int) {

	fmt.Println("El fibonacci de ", number, " Es: ")
	for n := 0; n < number; n++ {

		result := fibonacci(n)

		fmt.Println("Fibonacci ", n, " = ", result)
	}
}
func fibonacci(n int) int {
	a := 0
	b := 1

	if n < 0 {
		fmt.Print("Los fibonacci negativos no existen")
	} else {
		for i := 0; i < n; i++ {
			temp := a
			a = b
			b = temp + a
		}
	}
	return a
}

func factorial(n int) uint64 {
	var factVal uint64 = 1

	if n < 0 {
		fmt.Print("Los factoriales negativos no existen")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
			fmt.Println("Factorial ", i, " = ", factVal)
		}
	}
	return factVal
}
