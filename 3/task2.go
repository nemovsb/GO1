package main

import (
	"fmt"
	"math"
	"time"
)
func main(){
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	start := time.Now()
	fmt.Println(2)	// 2 - единственное четное простое число. Напечатаем его отдельно
						// (т.к. остаток от деления всегда будет = 0)
	for n := 3; n <= num; n += 2 {	// n - нечетные числа для проверки
		dmax := int(math.Sqrt(float64(n))) + 1	// dmax - максимально возможный потенциальный делитель
		for d := 2; d <= dmax; d++ {	// d - перебираем все возможные делители
			if n % d == 0 {	// если нет остатка от деления, то число нам не подходит
				break
			} else if d == dmax {	// если добрались до конца цикла, значит число простое
				//fmt.Println(n)
			}
		}
	}
	fmt.Printf("Elapsed time: %v ns\n", time.Since(start).Nanoseconds())
}
