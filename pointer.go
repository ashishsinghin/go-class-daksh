package main

// var p *int

// i := 42
// p = &i
// fmt.Println(*p)

import "fmt"

var i = 90

func main() {
	i, j := 42, 2701
	i = 56
	p := i

	fmt.Println(p)

	q := &j
	fmt.Println(q)
	fmt.Println(*q)
	fmt.Println(&q)
	fmt.Println(&j)
}
