package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i
	fmt.Println(*p) // Shows Value of pointer P which is looking at value I
	fmt.Println(p)  //This looks are the address of I

	*p = 99
	fmt.Println(i)

	p = &j
	*p = *p / 5
	fmt.Println(j)
}
