//method is the function with a receiver argument

//below is "Abs" function written as a regular function withnno change in fucntionality

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	//return math.Sqrt(x.X*v.X + v.Y*v.Y)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
