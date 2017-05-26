package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Place": Vertex{
		22, 33,
	},
	"Doodle": Vertex{
		1, 2,
	},
}

func main() {
	fmt.Println(m)
}
