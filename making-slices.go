package main

import "fmt"

func main() {
	a := make([]int, 5)
	fancySlicePrint("a", a)

	b := make([]int, 0, 5)
	fancySlicePrint("b", b)

	c := b[:2]
	fancySlicePrint("c", c)

	d := c[2:5]
	fancySlicePrint("d", d)
}

func fancySlicePrint(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
