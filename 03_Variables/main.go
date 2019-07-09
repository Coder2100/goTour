package main

import "fmt"

var c, python, lua, java bool

//bool for boolean datatype

//initialize variables
var i, j int = 1, 2

func main() {
	var a int
	var graphql, javascript, draftjs = true, false, "no!"
	fmt.Println(a, c, python, java)

	fmt.Println(i, j, javascript, graphql, draftjs)

	//short variable declaration
	var one, two int = 1, 2
	three := 3
	zukile, odwa, sam := true, false, "no!"

	fmt.Println(one, two, three, zukile, odwa, sam)

}
