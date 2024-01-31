package main

import "fmt"

const ebulicaoK = 373.2

func main() {

	k := ebulicaoK
	c := k - 273

	fmt.Printf("O valor do ponto de ebululição da água em Kelvin é: %g, já em Celsius é: %g.", k, c)

}
