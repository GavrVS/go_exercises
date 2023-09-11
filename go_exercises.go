package main

import (
	"fmt"
	"math"
)

func main() {

	//Exercise: Loops and Functions
	fmt.Println(Sqrt(16))
	fmt.Println(math.Sqrt(16))
}

func Sqrt(x float64) float64 {
	z := x / 2.0
	for (z*z - x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
