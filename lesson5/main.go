package main

import (
	"fmt"
	"os"
)

func main() {
	var num, i int32
	fmt.Print("Введите номер последовательности из рода Фибоначей:")
	_, err := fmt.Scanln(&num)
	if err != nil {
		os.Exit(1)
	}

	// fib1 := 0
	// fib2 := 1

	// for num > 0 {
	// 	fib1, fib2 = fib2, fib1+fib2
	// 	fmt.Println(fib1, ";", fib2, " = ", fib1+fib2)
	// 	num--
	// }

	fmt.Println(Fibo(num))

	var fib []int32
	for i = 0; i < num; i++ {
		fib = append(fib, 0)
	}

	FiboSlice(num-1, fib[:])
	fmt.Println(fib)
}

func Fibo(fnum int32) int32 {
	if fnum == 1 || fnum == 2 {
		return 1
	} else {
		return Fibo(fnum-1) + Fibo(fnum-2)
	}
}

func FiboSlice(fnum int32, array []int32) int32 {
	if fnum == 0 || fnum == 1 {
		array[fnum] = 1
	} else {
		array[fnum] = FiboSlice(fnum-2, array) + FiboSlice(fnum-1, array)
	}
	return array[fnum]
}
