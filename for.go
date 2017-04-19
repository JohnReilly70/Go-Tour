package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 20; i++ {
		sum += 1
		fmt.Printf("Sum this far %v \n", sum)
		fmt.Printf("It is type %T \n", sum)
	}

	fmt.Println(sum)
}
