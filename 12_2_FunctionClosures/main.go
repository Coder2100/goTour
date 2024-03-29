//Function closure

//A closure is a function value that references variables from outisde its body

//The function may access and assign to the referenced variables,
// in this sense the function is 'bound' to the variables

// eg the adder below function returns a closure.Each closure
// is bound to its own sum variable

package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
