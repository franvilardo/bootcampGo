package main

import "fmt"

func imprimirCantidadLetrasYLetrasPalabra(palabra string) {

	fmt.Println("La cantidad de letras es:", len(palabra))
	fmt.Println("Las letras son:")
	for _, letraSegundaVuelta := range palabra {
		fmt.Println(string(letraSegundaVuelta))
	}
}

func main() {
	imprimirCantidadLetrasYLetrasPalabra("total")
}
