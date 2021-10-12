package main

import (
	"fmt"
	"strconv"
	"strings"
)

// this is a comment

func main() {
	var cnt string
	fmt.Println("Введите массив чисел через *, %, _ или точку с запятой: ")
	fmt.Scanln(&cnt)

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%_;", r)
	}

	numstr := strings.FieldsFunc(cnt, splitFunc)

	var numb []int64

	for i := 0; i < len(numstr); i++ {
		tmp, err := strconv.ParseInt(numstr[i], 10, 64)
		if err == nil {
			numb = append(numb, tmp)
		}

	}

	Sort(numb)

	fmt.Println(numb)
}

func Sort(arr []int64) {

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}

	}
}
