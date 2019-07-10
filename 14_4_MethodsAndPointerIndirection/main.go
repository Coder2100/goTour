//The equivalent thing happens in the reverse direction.

//Functions that take a value argument must take a value of that specific type:

// var v Vertex

// fmt.Println(AbsFunc(v))// ok

//fmt.Println(AbsFunc(&v))//compile error

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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
