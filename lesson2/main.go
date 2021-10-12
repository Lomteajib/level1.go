package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float32

	fmt.Print("Length: ")
	fmt.Scanln(&a)

	fmt.Print("Height: ")
	fmt.Scanln(&b)

	fmt.Printf("Square: %f\n", a*b)

	//const pi = 3.1415926
	var csq float64

	fmt.Print("Circle square: ")
	fmt.Scanln(&csq)

	//csq = pi*r^2 => r=sqr(csq/pi)
	//csq = clen /4*pi => clen = sqr(4*pi*csq)

	fmt.Printf("Circle length: %f\n", math.Sqrt(4*math.Pi*csq))
	fmt.Printf("Circle radius: %f\n", math.Sqrt(csq/math.Pi))

	var i int
	fmt.Print("enter 3 digit number: ")
	fmt.Scanln(&i)

	fmt.Println("hundreds:", i/100)
	fmt.Println("decades:", i/10-(i/100)*10)
	fmt.Println("units:", i-(i/100)*100-10*(i/10-(i/100)*10))
}
