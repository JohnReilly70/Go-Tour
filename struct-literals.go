package main

import "fmt"

type CoOrds struct {
	X, Y int
}

var (
	v1 = CoOrds{1, 2}
	v2 = CoOrds{Y: 2}
	v3 = CoOrds{}
	p  = &CoOrds{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
