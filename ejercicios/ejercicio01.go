package ejercicios

import "strconv"

func ValidaNumero(cadena string) (int, string) {

	valorNumerico, errorConversion := strconv.Atoi(cadena)

	if errorConversion != nil {
		return 0, "El valor ingresado no es un número"
	}

	return valorNumerico, func() string {
		if valorNumerico > 100 {
			return "Es mayor a 100"
		}
		return "Es menor o igual a 100"
	}()

}
