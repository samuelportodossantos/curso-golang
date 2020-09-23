package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2 // ou exclusivo
	comprarSorver := trab1 || trab2

	return comprarTV50, comprarTV32, comprarSorver
}

func main() {
	tv50, tv32, sorvete := comprar(true, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
