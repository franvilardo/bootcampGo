/*

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
1. Categoría C, su salario es de $1.000 por hora.
2. Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
3. Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.


*/

package main

import "fmt"

func main() {
	fmt.Println(calcularSalario(4890, "A"))
}

func calcularSalario(cantidadMinutos int, categoria string) (salario float64) {
	switch categoria {
	case "A":
		salario = minutosAHoras(cantidadMinutos, 3000) * 1.5
	case "B":
		salario = minutosAHoras(cantidadMinutos, 1500) * 1.2
	case "C":
		salario = minutosAHoras(cantidadMinutos, 1000)
	}
	return salario
}

func minutosAHoras(minutos, valorHora int) (monto float64) {

	monto = float64(minutos) / float64(60) * float64(valorHora)

	return monto
}
