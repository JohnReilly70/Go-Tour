package main

import "fmt"

func main() {
	fmt.Println("STARTING DEFER LOOP")

	for i := 0; i < 11; i++ {
		defer fmt.Printf("%v", i)
	}

	fmt.Println("Finished")
}
