//interfaces are implemented implicitly

// a type implements an interface by implemnting its method

//there is no explicit declaration of intent, no 'implements' keywords

//implicit interefaes decouple the defintion of an interface from its implementation
// which could appear in any package without a prearrangement

package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implemts the interface I
// but we dont need to explicitly declare that it does

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
