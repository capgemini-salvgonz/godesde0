package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func OtrasVariables() {
	fmt.Println("Otras variables")
	Nombre = "Chava"
	Estado = true
	Sueldo = 1000.50
	Fecha = time.Now()

	fmt.Println("Nombre:", Nombre)
	fmt.Println("Estado:", Estado)
	fmt.Println("Sueldo:", Sueldo)
	fmt.Println("Fecha:", Fecha)
}

func ConvertirATexto(numero int) (bool, string) {
	var cadena = strconv.Itoa(numero)
	return true, cadena
}
