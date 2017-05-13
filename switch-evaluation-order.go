package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday() - 6
	fmt.Printf("%v\n", today)
	/*	switch time.Sunday {
		case today + 1:
			fmt.Println("Today")
		case today + 1:
			fmt.Println("Tomorow")
		case today + 2:
			fmt.Println("Two Days time!")
		case today + 3:
			fmt.Println("Three Days Time!")
		default:
			fmt.Println("AGES AWAY!!! :(")

		}*/
}
