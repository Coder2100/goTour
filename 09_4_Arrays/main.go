package main

import "fmt"

//SLICE DEFAULT

//When slicing you may omit  or low dounds to use their defaults instead

// the default is zero for the low bound and lenght of the slice for the high bound

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]

	fmt.Println(s)

	s = s[1:]

	fmt.Println(s)
}
