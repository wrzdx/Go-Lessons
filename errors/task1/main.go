package main

import (
	"errors"
	"fmt"
)

func checkNumbers(a int, b int) error {
	if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
		return errors.New("Arguments should be [-1000, 1000]")
	}

	return nil
}
func Add(a int, b int) (int, error) {
	if ok := checkNumbers(a, b); ok != nil {
		return 0, ok
	}

	return a + b, nil
}

func Subtract(a int, b int) (int, error) {
	if ok := checkNumbers(a, b); ok != nil {
		return 0, ok
	}

	return a - b, nil
}

func Multiply(a int, b int) (int, error) {
	if ok := checkNumbers(a, b); ok != nil {
		return 0, ok
	}

	return a * b, nil
}

func Divide(a int, b int) (int, error) {
	if ok := checkNumbers(a, b); ok != nil {
		return 0, ok
	}

	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}

func main() {
	examples := [][2]int{
		{1, 1},
		{1001, 1},
		{1, 0},
	}

	for _, pair := range examples {
		a, b := pair[0], pair[1]
		result, ok := Divide(a, b)
		if ok == nil {
			fmt.Println("Ok:", result)
		} else {
			fmt.Println("Operation: Division")
			fmt.Println("a =", a, ", b=", b)
			fmt.Println("Error:", ok.Error())
		}
	}
	
}
