package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	var x, y = 0, 1
	return func() (z int) {
		z, x, y = x, y, x+y
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
