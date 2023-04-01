package main

import (
	"fmt"

	operaciones "github.com/hectordmz/go_test/helpers"
	munBase "github.com/hectordmz/go_test/helpers/base"
)

func main() {

	text_menu := "Selecciones la opción de la operación que desea realizar\n"
	text_menu += "1 Suma\n"
	text_menu += "2 Resta\n"
	text_menu += "3 Multiplicación\n"
	text_menu += "4 División\n"
	text_menu += "5 Base de número\n"
	fmt.Println(text_menu)

	fmt.Print("Opción: ")
	var opc float64
	fmt.Scan(&opc)
	if opc != 5 {
		fmt.Print("Num1: ")
		var text float64
		fmt.Scan(&text)
		fmt.Print("Num2: ")
		var text2 float64
		fmt.Scan(&text2)
		nums := operaciones.Numeros{Numero1: text, Numero2: text2}
		fmt.Print("Respuesta: ", operaciones.Opcion(opc, nums))
	} else if opc == 5 {
		fmt.Print("Num: ")
		var num string
		fmt.Scan(&num)
		fmt.Print("Base: ")
		var base int
		fmt.Scan(&base)
		munBase.NumBase(munBase.NumerosBase{Num: num, Base: base})
	} else {
		fmt.Println("Opcion incorrecta")
	}
}
