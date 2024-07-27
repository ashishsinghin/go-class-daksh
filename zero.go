package main

// 0 for numeric types
// false for boolean types
// "" (the empty string) for strings

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
