package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	_, aerr := fmt.Scanln(&a)

	if aerr != nil {
		fmt.Printf("Введено некорректное число\n")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	_, berr := fmt.Scanln(&b)

	if berr != nil {
		fmt.Printf("Введено некорректное число\n")
		os.Exit(1)
	}
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)

}
