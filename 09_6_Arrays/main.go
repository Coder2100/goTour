package main

import "fmt"

//Nil slices

// a nill slcie has a length and capacity of nil

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil")
	}
}
