package main

import (
	"fmt"
	"runtime"
)

func main() {

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Linux")
	case "Windows":
		fmt.Println("Windows")
	default:
		fmt.Println("No es Linux ni Windows, es", os)
	}

	// files.WriteStringToFile(ejercicios.TablaDeMultiplicar(), "./files/txt/tabla.txt")
	// files.AppendStringToFile(ejercicios.TablaDeMultiplicar(), "./files/txt/tabla.txt")
	// files.AppendStringToFile(ejercicios.TablaDeMultiplicar(), "./files/txt/tabla.txt")
	// files.ReadFile("./files/txt/tabla.txt")

	// mac.GetRootBetween(0, 10, func(value float64) float64 { return math.Pow(value, 2) - 9 })

	/*
		funciones.Calculate()
		if funciones.IntPredicate(func(numero int) bool { return numero >= 10 }, 8) {
			fmt.Println("El numero evaluado es mayor o igual 10")
		} else {
			fmt.Println("El numero evaluado es menor a 10")
		}

		counter := funciones.CallsCounter()
		fmt.Printf("Llamadas a la función: [%d]\n", counter())
		fmt.Printf("Llamadas a la función: [%d]\n", counter())
		fmt.Printf("Llamadas a la función: [%d]\n", counter())
	*/

	/*
		user1 := repository.CreateUser(1, "Chava", "chava.gnolasco@gmail", true)
		user2 := repository.CreateUser(2, "Salvador", "salvador.gnolasco@gmail", false)

		userMap := make(map[string]models.User)

		userMap["user1"] = user1
		userMap["user2"] = user2

		userSlide := make([]models.User, 0)
		userSlide = append(userSlide, user1)

		fmt.Println("Users: ", userMap)
		fmt.Println("User Slide: ", userSlide[0:])
	*/

}
