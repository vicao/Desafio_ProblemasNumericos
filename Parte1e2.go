package main

import "fmt"

func main() {
	DesafioParte1()
	fmt.Println("================================")
	DesafioParte2()
}

func DesafioParte1() {
	fmt.Println("Desafio parte 1")
	/**
	Implementação do desafio parte 1 - Esse desafio terá como objetivo praticar estruturas de repetição em Go, uso do operador %.
	Voce e seus colegas querem criar um código que exiba todos os números entre 1 e 100 que são divisíveis por 3.
	*/
	fmt.Println("Números entre 1 e 100 divisíveis por 3:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Fim do desafio parte 1")
}

func DesafioParte2() {
	fmt.Println("Desafio parte 2")
	/**
	Implementação do desafio parte 2 - Esse desafio terá como objetivo praticar estruturas de repetição em Go, uso do operador %.
	Voce e seus colegas querem criar em formato de código uma antiga brincadeira:
	Ao falar numeros sempre que o numero for multiplo de 3 voce fala "Pin", se for multiplo de 5 voce fala "Pan".
	Então, seu programa deve imprimir os números de 1 a 100, substituindo os multiplos de 3 por "Pin", e os multiplos de 5 por "Pan".
	*/
	fmt.Println("Números de 1 a 100 com substituições:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}

}
