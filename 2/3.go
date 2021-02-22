package main

import "fmt"

func main(){
	var (
		num, hund, deci, unit uint32
		)

	fmt.Print("Введите трехзначное положительное число: ")
	fmt.Scanln(&num)
	if (num < 1000) {
		hund = num / 100
		deci = (num % 100) / 10
		unit = (num % 100) % 10
		fmt.Print("Сотни: ", hund, "\n")
		fmt.Print("Десятки: ", deci, "\n")
		fmt.Print("Единицы: ", unit, "\n")
	} else {
		fmt.Print("Число не должно быть больше 1000 \n")
	}

}
