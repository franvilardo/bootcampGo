/*

Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un
salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).

*/

package salario

// import "fmt"

// func main() {
// 	fmt.Println(calcularSalario(160000))
// }

func calcularSalario(sueldo float64) float64 {
	var respuesta float64
	if sueldo >= 50000 {
		respuesta = sueldo * 0.17
	}
	if sueldo >= 150000 {
		respuesta += sueldo * 0.1
	}
	return respuesta
}
