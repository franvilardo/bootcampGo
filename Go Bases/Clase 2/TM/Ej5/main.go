/*

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.


*/

package main

import "fmt"

func main() {
	funcionAUsar, mensaje := animal("Pato")
	if mensaje != "" {
		println(mensaje)
		return
	}
	fmt.Println(funcionAUsar(5))
}

func animal(cualAnimal string) (func(float64) float64, string) {
	switch cualAnimal {
	case "Perro":
		return calculoPerro, ""
	case "Gato":
		return calculoGato, ""
	case "Hamster":
		return calculoHamster, ""
	case "Tarántula":
		return calculoTarantula, ""
	default:
		return nil, "No tenemos de este animal. Esperamos poder darle refugio pronto."
	}
}

func calculoPerro(cantidadPerros float64) float64 {
	return cantidadPerros * 10
}

func calculoGato(cantidadGatos float64) float64 {
	return cantidadGatos * 5
}

func calculoHamster(cantidadHamsters float64) float64 {
	return cantidadHamsters * 0.25
}

func calculoTarantula(cantidadTarantulas float64) float64 {
	return cantidadTarantulas * 0.15
}
