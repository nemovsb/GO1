package main

import (
	"fmt"
	"math"
)

func main(){
	var squareCircle float64

	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&squareCircle)

	diameter, circumference := circleProperties(squareCircle)

	fmt.Print("Диаметр окружности: ", diameter, "\n")
	fmt.Print("Длина окружности: ", circumference, "\n")
}

func circleProperties(squareCircle float64) (float64, float64) {
	if (squareCircle >= 0) {
		var diameter, circumference float64

		diameter = math.Sqrt( squareCircle * 4 / math.Pi )
		circumference = math.Pi * diameter
		return diameter, circumference
	} else {
		panic("Отрицательная площадь.\n")
	}
}