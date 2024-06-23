package main

import (
	"fmt"
	"runtime"

	"github.com/chava.gnolasco/godesde0/ejercicios"
	"github.com/chava.gnolasco/godesde0/variables"
)

func main() {
	estado, nombre := variables.ConvertirATexto(64)
	fmt.Println("Estado:", estado)
	fmt.Println("Nombre:", nombre)

	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Linux")
	} else {
		fmt.Println("No es Linux, es", os)
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Linux")
	case "Windows":
		fmt.Println("Windows")
	default:
		fmt.Println("No es Linux ni Windows, es", os)
	}

	numero, mensaje := ejercicios.ValidaNumero("200")
	fmt.Println("Número:", numero)
	fmt.Println("Mensaje:", mensaje)

}
