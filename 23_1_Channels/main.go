package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to the c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	C := make(chan int)
	go sum(s[:len(s)/2], C)
	go sum(s[len(s)/2:], C)

	x, y := <-C, <-C // receive from c
	fmt.Println(x, y, x+y)
}
