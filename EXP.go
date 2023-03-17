package main

import "fmt"

func main() {

	fmt.Println(GetStatExp(17))
}

func GetStatExp(stat int) int {

	var experiencia = 0
	var cost = 0
	var decenas = stat / 10
	var unidades = stat % 10
	for k := 0; k < decenas; k++ {
		cost = k + 1
		for i := 0; i < 10; i++ {
			experiencia = experiencia + cost
		}
	}
	return experiencia + (unidades * (cost + 1))
}
