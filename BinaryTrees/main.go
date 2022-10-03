package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		_walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for v1 := range c1 {
		v2 := <-c2
		if v1 != v2 {
			fmt.Println(v1, v2)
			return false
		}
	}
	// making sure no more elements in both trees
	_, ok1 := <-c1
	_, ok2 := <-c2
	return !ok1 && !ok2
}

func main() {
	c := make(chan int)
	t := tree.New(10)
	t1 := tree.New(10)

	go func() {
		// for i := 0; i < 10; i++ {
		// 	fmt.Println(<-c)
		// }
		//receives values from the channel repeatedly until it is closed.
		for v := range c {
			fmt.Println(v)
		}
	}()
	Walk(t, c)
	if Same(t1, t) {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}
