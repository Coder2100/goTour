package main

import (
	"fmt"
	"math"
)

//IFstatements are like for loops, there no need for () but the {} are required
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//

//IF with a shorthand
//POWER
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

//IF ELSE
//variables declared inside if short statement are also available inside any of the else blocks

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// the init and post statement are optional

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	//repeat part 1 not ideal as it is consufssing

	sum1 := 1
	for sum1 < 10 {
		sum1 += sum1
	}

	fmt.Println(sum1)

	// For is Go's 'while'
	//repeat of of sum2
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}

	fmt.Println(sum3)

	//Forever

	// if you ommit the loop condition if loops forever
	// infinite loop is compactly expressed

	// using sqrt
	fmt.Println(sqrt(2), sqrt(-4))

	//POWER FUNCTION

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
