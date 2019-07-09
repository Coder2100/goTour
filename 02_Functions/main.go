package main

import "fmt"

func add(x int, y int) int { // data type comes after the variable
	return x + y
}

// declaring variables short

func subtract(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //it returns the declared variables
}
func main() {
	fmt.Println(add(42, 13))
	fmt.Println(subtract(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
