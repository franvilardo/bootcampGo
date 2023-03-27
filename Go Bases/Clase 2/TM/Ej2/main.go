/*

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/

package main

import "fmt"

func main() {
	fmt.Println(promedioNotas(5, 6, 8, 10, 4, 8))
}

func promedioNotas(notas ...uint) (average float64) {
	for _, nota := range notas {
		average += float64(nota)
	}
	average = average / float64(len(notas))
	return average
}
