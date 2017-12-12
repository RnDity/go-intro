package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	for i := 0; i < 10; i++ {
		fmt.Println(i, x, y, k, c, python, java)
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

