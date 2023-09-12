package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
	"strings"
)

func main() {

	//Exercise 1: Loops and Functions
	number := 16.0
	fmt.Printf("Result sqrt:\n custom: %f\n math: %f\n", Sqrt(number), math.Sqrt(number))

	//Exercise 2: Slices
	pic.Show(Pic)

	//Exercise 3: Maps
	wc.Test(WordCount)

	//Exercise 4: Fibonacci closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func Sqrt(x float64) float64 {
	z := x / 2.0
	for (z*z - x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}
		picture[y] = row
	}
	return picture
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCountMap := make(map[string]int)
	for _, word := range words {
		count, ok := wordCountMap[word]
		if !ok {
			wordCountMap[word] = 1
		} else {
			wordCountMap[word] = count + 1
		}
	}
	return wordCountMap
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}
