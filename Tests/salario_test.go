package salario

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularSalario(t *testing.T) {
	t.Run("Salario por debajo de $50000", func(t *testing.T) {
		//Arrange
		var salarioIngresado float64 = 40000.0
		var impuestoEsperado float64 = 0.0

		//Act
		impuestoObtenido := calcularSalario(salarioIngresado)

		//Assert
		assert.Equal(t, impuestoEsperado, impuestoObtenido, "el impuesto obtenido no es el esperado")
	})
	t.Run("Salario entre $50000 y $150000", func(t *testing.T) {
		//Arrange
		var salarioIngresado float64 = 60000.0
		var impuestoEsperado float64 = 10200.0

		//Act
		impuestoObtenido := calcularSalario(salarioIngresado)

		//Assert
		assert.Equal(t, impuestoEsperado, impuestoObtenido, "el impuesto obtenido no es el esperado")
	})
	t.Run("Salario mayor a $150000", func(t *testing.T) {
		//Arrange
		var salarioIngresado float64 = 160000.0
		var impuestoEsperado float64 = 43200.0

		//Act
		impuestoObtenido := calcularSalario(salarioIngresado)

		//Assert
		assert.Equal(t, impuestoEsperado, impuestoObtenido, "el impuesto obtenido no es el esperado")
	})

}
