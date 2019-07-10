//comparing the previous two pograms , you will notice nthat,
//the functions with a pointer argument must take a pointer

//var v Vertex

//ScaleFunc(v 5)// compiles error!

//ScaleFunc(&v, 5)//ok

//while methods with pointer receivers take either a value or a pointer,
//as a receiver when they called

//for the statement v.Scale(5), even though v is value and not a
//pointer, the method with the pointer receiver is called automatically.package main

//go inteprete v.Scale(5) as (&v).Scale(5) since the scale method has a pointer receiver

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
