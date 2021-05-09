package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Print(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
