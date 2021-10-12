package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a int
	fmt.Print("введите число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Print("введено некорректное число")
		os.Exit(1)
	}
again:
	for i := 3; i <= a; i += 2 {
		m := int(math.Floor(math.Sqrt(float64(i))))
		for j := 2; j <= m; j++ {
			if i%j == 0 {
				continue again
			}
		}
		fmt.Println(" ", i)
	}

}
