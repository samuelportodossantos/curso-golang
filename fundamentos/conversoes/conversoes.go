package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9

	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//cuidado na hora de concatenar
	fmt.Println("testee " + string(97))

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string para int
	num, erro := strconv.Atoi("123321")

	fmt.Println(num)
	fmt.Println(erro)

}
