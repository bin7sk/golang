package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	x, y int
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.x*v.x + v.y*v.y))
}

func main() {
	v := Vertex{4,3}
	fmt.Println(v.Abs())
}
