// A map maps keys to values

//The zero value of a map is nil

//a nil map has no key or can keys be added

//The make function return a map of given type, initialized and ready for use

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.3967,
	}
	fmt.Println(m["Bell Labs"])
}
