package main

import (
	"fmt"
)

func firstelement[T any](a []T) T {
	return a[0]
}

func main() {
	intslice := []int{1, 2, 3}
	fmt.Println(firstelement(intslice))

	stringslice := []string{"ashish", "daksh", "rahul"}
	fmt.Println(firstelement(stringslice))

	floatslice := []float64{10.2, 30.3, 20.4}
	fmt.Println(firstelement(floatslice))
}
