package main

import "fmt"

func main() {

	i := 1

	//Go não tem arimetica de ponteiros
	var p *int = nil

	p = &i // pegando endereço de i e atribuindo ao ponteiro

	*p++ // Desreferenciando ponteiro pra pegar o valor
	i++

	fmt.Println(*p, p, i)
}
