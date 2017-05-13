package main

import "fmt"

func main() {

	names := [4]string{
		"John",
		"Paul",
		"Stevo",
		"Jimbo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]
	fmt.Println(a, b)

	b[0] = "X Gon Give It To Ya!"
	fmt.Println(a, b)
	fmt.Println(names)

}
