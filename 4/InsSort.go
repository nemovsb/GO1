package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	arr := []int64{}
	fmt.Print("Enter the elements: \n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
	}
	start := time.Now()
	InsSort(arr)
	fmt.Println(arr)
	fmt.Printf("Elapsed time: %v ns\n", time.Since(start).Nanoseconds())
}

func InsSort(arr []int64) {
	for i, _ := range arr { // Перебираем элементы несортированной части среза
		for j := i; j >= 1 && arr[j-1] > arr[j]; j-- {
			// Если число справа больше числа слева выходим из цикла
			arr[j], arr[j-1] = arr[j-1], arr[j] // Переставляем местами элементы в сортированной части
			//fmt.Println(arr, "i = ", i, " j = ", j)
		}
	}
}
