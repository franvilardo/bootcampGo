package main

import "fmt"

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func main() {
	var salary int = 140000

	if salary < 150000 {
		myErr := &myError{message: "el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(myErr)

	} else {
		fmt.Println("Debe pagar impuestos")
	}

}
