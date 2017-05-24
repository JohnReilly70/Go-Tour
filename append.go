package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3) // When Appened to a list be aware the Capacity may become larger than the length of the list. Makes more space if lots of items appeneded to a list
	printSlice(s)
}

func printSlice(list []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(list), cap(list), list)
}
