package main

import "fmt"

func main() {

	x := 1
	y := 2

	// apenas postfix
	x++
	fmt.Println(x)

	y--
	fmt.Println(y)

	// fmt.Println(++x == y--) não pode incrementar ou decrementar dentro de uma função

}
