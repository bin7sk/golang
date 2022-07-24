package main

import "fmt"

type Vertex struct {
	x, y int
}

var m map[string]Vertex
func main() {
	m = make(map[string]Vertex)
	m["key"] = Vertex{1, 2}
	m["dot"] = Vertex{0, 0}
	fmt.Println(m)
}
