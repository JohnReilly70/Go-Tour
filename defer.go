package main

import "fmt"

func main() {
	defer fmt.Println("DEFERED LINE")

	fmt.Println("Normal Print\n All Lines will run before defer lines will be called\nLets watch this print defered line after this long winded print")

}
