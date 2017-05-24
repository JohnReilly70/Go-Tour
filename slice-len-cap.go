package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	s = s[:0] //gives slice zero length
	printSlices(s)

	s = s[:4] // ented its length by 4
	printSlices(s)

	s = s[2:] // drop first 2 values
	printSlices(s)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
