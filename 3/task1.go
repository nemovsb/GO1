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
	fmt.Scanln(&a)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, №): ")
	fmt.Scanln(&op)

	if op != "№" {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}


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
	case "№": {
		if (a < 0) {
			fmt.Println("Искомое число - комплексное")
			os.Exit(1)
		} else {
			res = math.Sqrt(a)
		}
	}
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}