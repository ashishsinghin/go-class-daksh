package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("Ashish")
	fmt.Println("Hello")
}
