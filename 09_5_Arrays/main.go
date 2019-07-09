package main

import "fmt"

//Slice length and capacity

//A slice has both a length and a capacity.

//The length and
//capacity of a slice s can be obtained using the expressions len(s) and cap(s).

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//slice the slice to give it zero length

	s = s[:0]

	printSlice(s)

	// extends its length

	s = s[:4]
	printSlice(s)

	//drop its first two valus
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
