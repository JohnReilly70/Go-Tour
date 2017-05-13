package main

import "fmt"

type CoOrds struct {
	X int
	Y int
}

func main() {
	v := CoOrds{10, 20}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
