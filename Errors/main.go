package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	r, z := 1.0, x/2
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Sprint("the value is i, z, r", z, r)
		if math.Abs(z-r) < 0.000001 {
			return z, nil
		}
		r = z
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
