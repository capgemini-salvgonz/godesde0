package funciones

/**
 * CallsCounter returns a function that returns the number of times it has been called.
 *
 * It is an example of a closure, a function that captures the environment in which it was created.
 */
func CallsCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}
