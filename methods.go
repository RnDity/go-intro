package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Sum() float64 {
	return v.X + v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Sum())
}
