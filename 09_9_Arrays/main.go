//Appending to a slice

//appending new elements to a slice, suing append built in function

//the first parameter s of append is a slice of type T and the ret are T values ao append to the slice

package main

import "fmt"

//CREATE printSlice function

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)

	// append grows as need
	s = append(s, 1)

	printSlice(s)

	// we can add more than one element a time

	s = append(s, 2, 3, 4)

	printSlice(s)

}
