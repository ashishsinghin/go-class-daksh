package main

import "fmt"

type Vertex struct {
	X int
	Y int
	s string
}

func main() {
	v := Vertex{1, 2, "Ashish"}
	v.X = 4
	v.Y = 10
	fmt.Println(v)
	fmt.Println(v.X)
}
