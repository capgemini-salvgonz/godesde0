package funciones

import "fmt"

func Calculate() {

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(suma(10, 25))
}

func IntPredicate(fn func(int) bool, numberToEvaluate int) bool {
	return fn(numberToEvaluate)
}
