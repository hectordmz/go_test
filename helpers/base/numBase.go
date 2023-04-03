package munBase

import (
	"fmt"
	"math"
	"strconv"
)

type NumerosBase struct {
	Num  string
	Base int
}

func NumBase(aNumeros NumerosBase) {
	num := aNumeros.Num
	base := aNumeros.Base
	var newBase int

	decimal := convertToDecimal(num, base)
	fmt.Printf("El número %s en base %d es igual a %d en base 10\n", num, base, decimal)

	for {
		var answer string
		fmt.Println("¿Desea cambiar la base? (s/n)")
		fmt.Scan(&answer)

		if answer == "s" {
			fmt.Println("Ingrese la nueva base:")
			fmt.Scan(&newBase)

			newNum := convertFromDecimal(decimal, newBase)
			fmt.Printf("El número %d en base %d es igual a %s en base %d\n", decimal, newBase, newNum, newBase)
		} else if answer == "n" {
			fmt.Println("Saliendo...")
			break
		} else {
			fmt.Println("Respuesta no válida. Por favor, ingrese 's' o 'n'.")
		}
	}
}

func convertToDecimal(num string, base int) int {
	var decimal int
	for i := len(num) - 1; i >= 0; i-- {
		digit := num[i]
		var value int
		if digit >= '0' && digit <= '9' {
			value = int(digit - '0')
		} else {
			value = int(digit-'A') + 10
		}
		decimal += value * int(math.Pow(float64(base), float64(len(num)-i-1)))
	}
	return decimal
}

func convertFromDecimal(decimal, base int) string {
	var result string
	for decimal > 0 {
		remainder := decimal % base
		var digit string
		if remainder < 10 {
			digit = strconv.Itoa(remainder)
		} else {
			digit = string('A' + (remainder - 10))
		}
		result = digit + result
		decimal /= base
	}
	return result
}
