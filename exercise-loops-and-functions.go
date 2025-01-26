package main

import "fmt"

func Sqrt(x float64) float64 {
	// A decent starting guess for z is 1, no matter what the input.
	z := 1.0
	// To begin with, repeat the calculation 10 times and print each z along the way.
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		x := float64(i + i)
		fmt.Println(Sqrt(x))
	}
}
