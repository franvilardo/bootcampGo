// Crear un programa que cumpla los siguiente puntos:
// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
// Tener un slice global de Product llamado Products instanciado con valores.
// 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
// Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
// Ejecutar al menos una vez cada método y función definido desde main().

package main

import "fmt"

func main() {
	nuevoQueso := Product{
		ID:          1,
		Name:        "Queso Muzzarela",
		Price:       50.0,
		Description: "Queso graso que se derrite fácilmente",
		Category:    "Lácteos",
	}

	fmt.Println("******")
	nuevoQueso.GetAll()
	fmt.Println("******")
	nuevoQueso.Save()
	nuevoQueso.GetAll()
	fmt.Println("******")
	fmt.Println(getById(1))
	fmt.Println("******")
	fmt.Println(getById(2))
	fmt.Println("******")
	nuevoQueso.GetAll()

}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, eachProduct := range Products {
		fmt.Println(eachProduct)
	}
}

func getById(idToSearch int) Product {
	for _, eachProduct := range Products {
		if eachProduct.ID == idToSearch {
			return eachProduct
		}
	}
	return Product{
		ID:          -1,
		Name:        "Producto inexistente",
		Price:       0.0,
		Description: "No existe el producto. El ID ingresado no coincide",
		Category:    "Inexistente",
	}
}
