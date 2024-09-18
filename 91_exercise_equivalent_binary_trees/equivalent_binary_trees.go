// this exercise needs to be tested on tour playground

package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
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
	ch_1 := make(chan int)
	ch_2 := make(chan int)
	
	go Walk(t1, ch_1)
	go Walk(t2, ch_2)
	
	for i := 0; i < 10; i++ {
		if (<- ch_1) != (<- ch_2) {
			return false
		}
	}
	
	return true
}

func main() {
	ch := make(chan int)
	
	go Walk(tree.New(1), ch)
	
	for i := 0; i < 10; i++ {
		v := <- ch 
		fmt.Printf("%v ", v)
	}
	
	fmt.Println()
	
	fmt.Printf("Are trees equivalent? - %t\n", Same(tree.New(1), tree.New(1)))
	
	fmt.Printf("Are trees equivalent? - %t\n", Same(tree.New(1), tree.New(2)))
}
