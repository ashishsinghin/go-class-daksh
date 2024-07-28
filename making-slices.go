package main

import "fmt"

func main() {
	a := make([]int, 5) // var a []int{0,0,0,0,0} len = 5 cap = 5
	printSlice("a", a)

	b := make([]int, 1, 5) // var b []int{0} len = 1 cap = 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
