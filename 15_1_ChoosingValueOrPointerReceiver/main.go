//Tweo reasons to use a pointer receiver

// -1- the method can modify the value that its reiver points to

// -1- to avoid copying the values on each method call.package main

//this  can be more efficient if the receiver is large struct
//eg in the example below both scale abd abs are with the receiver type

//*Vertex, even though the Abs method needn't modify its receiver.

//In general, all methods on a given type should have either value or
// pointer receivers, but not a mixture of both.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	//fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
