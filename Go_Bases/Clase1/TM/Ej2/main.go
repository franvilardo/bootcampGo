//Ejercicio 2 - Clima

//Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.
//Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
//Imprimí los valores de las variables en consola.
//¿Qué tipo de dato le asignarías a las variables?

package main

import "fmt"

//temp float, humedad int, presion float
var temperatura, humedad, presion = 25.0, 70, 1023.8

func main() {
	fmt.Println("Temperatura:", temperatura, "\n", "Humedad:", humedad, "% \n", "Presión:", presion)
}
