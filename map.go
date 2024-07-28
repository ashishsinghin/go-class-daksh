package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["google"] = Vertex{40.23465, -12.2364}
	m["india"] = Vertex{23.2354, -23.2364}
	fmt.Println(m["google"])
	fmt.Println(m["india"])
}
