package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%v %v %v\n", s, len(s), cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("%v %v %v\n", s, len(s), cap(s))
	}
}
