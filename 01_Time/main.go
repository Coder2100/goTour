package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("The time is", time.Now())
	fmt.Println("My random number id: ", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//fmt.Println(math.pi) //undifined as pi is not explicitly imported
}
