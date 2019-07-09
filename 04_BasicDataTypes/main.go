package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//should use int unless you have specifi reason for integer values declaration
var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1

	z complex128 = cmplx.Sqrt(-5 + 12i)
)

//Numeric Constants
const (
	//create huge number by shifting a 1 bit left 100 places.
	//the binary number is 1 followed by 100 zeroes

	Big = 1 << 100
	//shift it right again 99 places, so we end up with 1 << 1, or 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero values or uninitialised variables are given zero values

	var zeroint int
	var zerofloat float64
	var zerobool bool
	var zerostring string

	fmt.Println(zeroint, zerofloat, zerobool, zerostring)
	//Type conversation
	// T(v) converts v to type T
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)

	fmt.Println(x, y, z)
	// Type Interface
	//if type unspecified, by using := syntx of var = expression,
	//the type is inferred from the value of on the right hand side
	var i int
	j := i //j is int
	//when the right hand side contains an untyped numeric constant,
	// the new variable may be an int, float64 or complex128
	//depending on the precision of the constant

	ii := 42           // int
	ff := 3.142        // float
	gg := 0.867 + 0.5i //complex128

	vv := 42 // change me!
	fmt.Println(vv, j, ii, gg, ff)
	fmt.Printf("v is of type %T\n", vv)

	//
	const Pi = 3.14
	//Constants
	//func main(){
	const World = "世界"

	fmt.Println("Hello", World)

	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}
