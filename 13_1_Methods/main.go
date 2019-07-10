// Go does not have classes
//method are descrined on types

//method is a function with a special receiver argument

//the receiver appears in its own argument list between th func
//key word and the method name

// below, the 'Abs' method has a receiver of type Vertex named v

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//mehtod
func (v Vertex) Abs() float64 { // abs is the receiver of type Vetex named v
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
