package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

func Sqrt(x float64) float64 {
	r, z := 1.0, x/2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("the value is i, z, r", i, z, r)
		if math.Abs(z-r) < 0.000001 {
			break
		}
		r = z
	}
	return z
}
