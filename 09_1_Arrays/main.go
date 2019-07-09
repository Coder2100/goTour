package main

import "fmt"

func main() {
	//declares an array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1]) // accessing an array
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//SLICES
	//an array has a fixed size.slice is flexible
	// the []T is a slice with elements of type T
	//A Slice is formed by specifying two indics, a low and high bound separated comma

	// a[low :high]

	primes2 := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes2[1:4]

	fmt.Println(s)

	// Slices are like references to arrays
	// a slice does not store any data, it just describes a section of an underlying array

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

}
