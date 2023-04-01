package operaciones

import "fmt"

type Numeros struct {
	Numero1 float64
	Numero2 float64
}

func Opcion(opcion float64, aNumeros Numeros) float64 {
	switch opcion {
	case 1:
		return Suma(aNumeros)
	case 2:
		return Resta(aNumeros)
	case 3:
		return Divicion(aNumeros)
	case 4:
		return Multiplicacion(aNumeros)
	case 5:
		return Suma(aNumeros)
	default:
		fmt.Println("Opcion Incorrecta !!")
		return 0
	}
}

func Suma(aNumeros Numeros) float64 {
	return aNumeros.Numero1 + aNumeros.Numero2
}
func Resta(aNumeros Numeros) float64 {
	return aNumeros.Numero1 - aNumeros.Numero2
}
func Divicion(aNumeros Numeros) float64 {
	return aNumeros.Numero1 / aNumeros.Numero2
}
func Multiplicacion(aNumeros Numeros) float64 {
	return aNumeros.Numero1 * aNumeros.Numero2
}
