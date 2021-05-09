package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

//tree structure
// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var treeWalk func(t *tree.Tree, ch chan int)
	treeWalk = func(t *tree.Tree, ch chan int) {
		if t.Left != nil {
			treeWalk(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			treeWalk(t.Right, ch)
		}
	}
	treeWalk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("True case:", Same(tree.New(1), tree.New(1)))
	fmt.Println("False case:", Same(tree.New(1), tree.New(2)))
}
