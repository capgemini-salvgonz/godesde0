package mac

/**
 * Función que implementa el método de bisección para encontrar la raíz de una función
 * en un intervalo [a, b] donde a es menor que b.
 */
func GetRootBetween(a, b float64, fn func(float64) float64) float64 {

	look := true
	var valA, valB float64 = a, b

	for look {
		left := fn(valA)
		right := fn(valB)

		look = !((left * right) == 0)

		if !look {
			return func() float64 {
				if left == 0 {
					return valA
				}
				return valB
			}()
		}

		if left*right < 0 {
			valA = (valA + valB) / 2
		} else {
			if left < 0 && right < 0 {
				valB += 1
			} else {
				valA -= 1
			}
		}

	}

	return (valA + valB) / 2
}
