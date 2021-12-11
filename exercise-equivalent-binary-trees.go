package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Part 1 of the exercise
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Part 3 of the exercise
	var t1Array [10]int
	var t2Array [10]int
	ch := make(chan int)
	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		t1Array[i] = <-ch
	}
	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		t2Array[i] = <-ch
	}
	return t1Array == t2Array
}

func main() {
	// Part 2 of the exercise
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// Part 4 of the exercise
	fmt.Println("Expected: true, Actual: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Expected: false, Actual: ", Same(tree.New(1), tree.New(2)))
}
