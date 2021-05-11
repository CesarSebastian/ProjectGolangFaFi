package main

import "fmt"

func main() {

	var numero int

	menu := "Ingresa el número que quieres transformar "

	fmt.Print(menu)

	fmt.Scanln(&numero)

	optionFacto(numero)  //Terminado
	optionFibona(numero) //Terminado
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
	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
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
