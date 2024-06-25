package ejercicios

import (
	"fmt"
	"strconv"
)

func TablaDeMultiplicar() string {

	var input string
	var number int
	var isError error
	var textResult string

	for readValue := true; readValue; {
		fmt.Print("Indique el valor para la tabla de multiplicar: ")
		fmt.Scanln(&input)
		number, isError = strconv.Atoi(input)

		if readValue = isError != nil; readValue {
			fmt.Println("\nEl valor ingresado no es un número, intente nuevamente")
		}
	}

	for i := 1; i <= 10; i++ {
		textResult += fmt.Sprintf("%2d x %2d = %3d\n", number, i, number*i)
	}

	return textResult
}
