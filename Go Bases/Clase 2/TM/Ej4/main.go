/*

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje (en caso que el cálculo no esté
definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior.


*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	calificaciones := []float64{2.0, 3.1, 5.2, 7.0, 9.5, 10, 12}
	sort.Slice(calificaciones, func(i, j int) bool {
		return calificaciones[i] < calificaciones[j]
	})

	funcionAUsar, mensaje := operacion("promedio")
	if mensaje != "" {
		fmt.Println(mensaje)
		return
	}
	fmt.Println(funcionAUsar(calificaciones))
}

func minimo(calificaciones []float64) float64 {
	return calificaciones[0]
}

func maximo(calificaciones []float64) float64 {
	return calificaciones[len(calificaciones)-1]
}

func promedioNotas(calificaciones []float64) (average float64) {
	for _, nota := range calificaciones {
		average += float64(nota)
	}
	average = average / float64(len(calificaciones))
	return average
}

func operacion(tipoOperacion string) (func([]float64) float64, string) {
	switch tipoOperacion {
	case "minimo":
		return minimo, ""
	case "maximo":
		return maximo, ""
	case "promedio":
		return promedioNotas, ""
	default:
		return nil, "operación no definida"
	}
}
