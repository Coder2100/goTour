//mutating maps
//insert or updates an elements in map m:

//m[key] = elem or value

//Retrieve an element

// elem = m[key]

//deleting an element

//delete(m, key)

//test that a key is present with a two-value assignment

//elem, ok = m[key]

// if key is in m, ok is true. if not ok, ok is false

// if key is in not in map, then elem si the zero value for the maps' element type

//NOTE, if elem or ok have not been yet declared you could use a shorthand form

// elem, ok := m[key]
// m for map

package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])
	//changin the exisiting value
	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	//deleting the value of the map
	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	//check if the key is present with two values assignment
	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present", ok)
}
