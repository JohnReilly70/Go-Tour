package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Place": {
		22, 33,
	},
	"Doodle": {
		1, 2,
	},
}

func main() {
	fmt.Println(m)
}
