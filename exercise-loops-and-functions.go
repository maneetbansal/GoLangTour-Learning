package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Try assigning different values to Z here
	z := x / 2
	delta := 0.001
	for i := 0; i < 10; i++ {
		y := (z - (z*z-x)/(2*z))
		fmt.Println("root is:", y)
		if math.Abs(y-z) < delta {
			fmt.Println("Stopping after:", i, "iterations")
			break
		}
		z = y
	}
	return z
}

func main() {
  fmt.Println(Sqrt(2))
	// Try Playing around with different values
  //fmt.Println(Sqrt(5))
}
