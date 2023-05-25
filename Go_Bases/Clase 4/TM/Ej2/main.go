package main

import (
	"fmt"
	"errors"
)

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func main () {
	var salary int = 8000

	if salary < 10000 {
		myErr := &myError{message: "el salario ingresado no alcanza el mínimo imponible"}
		var err *myError
		if errors.Is(myErr, err) {
			fmt.Println(myErr.message)
		}
type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func validate(salary int)error {
	if salary < 10000 {
		return errors.New("Error: el salario es menor a 10.000")
	}
	return nil
}

func main () {
	var salary int = 8000
	myErr := myErr.Error("")
	if condition {
		
	}
}