//POINTER RECEIVERS

//You can declare methods with a pointer receivers

//this means the receiver type has the literal syntax *T,
//for some type t
// ALSO, T cannot itself be a pointer such as *int

// for example, the Scale mehtod here is degine *Vertex

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
