package main

import "fmt"

func main() {
	fmt.Println("STARTING DEFER LOOP")

	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
		defer fmt.Printf("\nf%v	%T\n", i, i)
	}

	fmt.Println("Finished")
}
