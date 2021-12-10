package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// Uncommenting this line will put the code in infinte loop
	// Why? Because Go will try to call Error() method to convert
	// it to string.

	//fmt.Sprint(e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		// I used the sqrt definition created in earlier exercise
		// You can also use math.Sqrt(x)
		return sqrtPositive(x), nil
	}
}

// Copied from in earlier excercise
func sqrtPositive(x float64) float64 {
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
	fmt.Println(Sqrt(-2))
}
