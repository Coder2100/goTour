package main

import "fmt"

func main() {
	i, j := 43, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j

	*p = *p / 37

	fmt.Println(j)
}

// the & operator genrate a pointer to its operand
// the * operator denotes the pointers underling value
// go does not have pointer airtmatic
