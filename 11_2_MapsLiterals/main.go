//map literals are like strut=ct literals, but the keys are required

package main

import "fmt"

type Vertex struct {
	lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.86433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.0840,
	},
}

func main() {
	fmt.Println(m)
}
