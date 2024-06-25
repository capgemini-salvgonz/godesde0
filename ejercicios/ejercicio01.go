package ejercicios

import "strconv"

func ValidaNumero(cadena string) (bool, int, string) {

	valorNumerico, errorConversion := strconv.Atoi(cadena)

	if errorConversion != nil {
		return false, 0, "El valor ingresado no es un número " + errorConversion.Error()
	}

	return true, valorNumerico, func() string {
		if valorNumerico > 100 {
			return "Es mayor a 100"
		}
		return "Es menor o igual a 100"
	}()

}
