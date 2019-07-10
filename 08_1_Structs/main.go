package main

import "fmt"

//struckt is the collection of fields
//part 1
type Vertex struct {
	X int
	Y int
}

//part 2
type Vertical struct {
	X1 int
	Y1 int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// struct filed are accessed using dot
	v := Vertical{1, 2}
	v.X1 = 4
	v.Y1 = 4
	fmt.Println(v.X1)
	fmt.Println(v.Y1)
	//POINTERS TO STRUCT
	//struct fields can be accessed through a struct pointer

	//use p.X to access struct
	//part 3
	p := &v
	p.X1 = 1e9
	fmt.Println(v)

}
