package main

import (
	"fmt"

	"github.com/chava.gnolasco/godesde0/variables"
)

func main() {
	estado, nombre := variables.ConvertirATexto(64)
	fmt.Println("Estado:", estado)
	fmt.Println("Nombre:", nombre)
}
