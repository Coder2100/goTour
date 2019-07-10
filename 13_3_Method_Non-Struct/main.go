//Declare a method on non-struct types

//numeric type MyFloar with an Abs method

//you can declare a method with a receiver whose type is ...package main

// defined in another package (which is the built in types such as int)
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
