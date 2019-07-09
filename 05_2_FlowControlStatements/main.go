package main

import (
	"fmt"
	"math"
)

// IF ELSE statements

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//cant use v
	return lim
}

//create SQUARE ROOT FUNCTION

func Sqrt(x float64) float64 {

	return x
}

//main function

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(2))
}
