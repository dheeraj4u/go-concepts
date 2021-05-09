package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
	// a := [3]int{1, 2, 3}
	// q := &a[1]
	// fmt.Println(*q)
}
