package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go run on")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		//freesbd opensbd
		//plan9, win32
		// nacl
		fmt.Printf("%s. \n", os)

	}
	fmt.Println(runtime.GOOS)

	//switch statement on weekdays
	fmt.Println("When's Friday?")
	today := time.Now().Weekday()
	switch time.Friday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//DEFER
	//A defer statement defers the execution of a function unitl the surrounding function returns

	// the deferred calls argument are evaluated immediately,
	//but the function call is not executed until the surrounding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	//STACKING DEFERS
	//Deferred function calls are pushed ont a STACK
	//when a function returns, its deferred calls are executed in last in first out order

	//
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
