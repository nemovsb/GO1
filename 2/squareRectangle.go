package main

import "fmt"

func main(){
	var lenghtRectangle, widthRectangle float64

	fmt.Print("Введите длину прямоугольника: ")
	fmt.Scanln(&lenghtRectangle)

	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&widthRectangle)

	fmt.Print("Площадь прямоугольника: ", squareRectangle(lenghtRectangle, widthRectangle),"\n")
}

func squareRectangle(lenghtRectangle, widthRectangle float64) float64 {
	return lenghtRectangle * widthRectangle
}
