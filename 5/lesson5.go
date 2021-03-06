package main

import "fmt"

type fibonacci map[uint64]uint64

func main(){
	var n uint64
	var fib fibonacci = map[uint64]uint64{1:0, 2:1}

	fmt.Print("Enter the Fibonacci number index: ")
	fmt.Scanln(&n)
	fmt.Println("Number: ",fib.calcNumber(n))

	fmt.Print("Enter another Fibonacci number index: ")
	fmt.Scanln(&n)
	fmt.Println("Number: ",fib.calcNumber(n))
}

func (f *fibonacci) calcNumber(n uint64) uint64 {
	_, exist := (*f)[n]	//Есть ли такое число во входной мапе?
	fmt.Println("Need a calculation? : ",!exist)
	if exist {	//Есть.
		return (*f)[n]
	} else {  //Нет такого числа, придется считать
		for i := uint64(len(*f)) + 1; i <= n; i++ {
			//fmt.Println("i = ", i)
			(*f)[i] = (*f)[i-2] + (*f)[i-1]
		}
		fmt.Println(*f)
		return (*f)[n]
	}
}