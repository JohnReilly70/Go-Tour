package main

import "fmt"

type Co_Ords struct {
	X float64
	Y float64
	Z float64
}

func main() {
	point := Co_Ords{1, 2, 3}
	fmt.Println(Co_Ords{3.14, 2.71828, 1.41421})
	fmt.Println(point)
	fmt.Printf("X:%v \nY:%v \nZ:%v \n", point.X, point.Y, point.Z)
}
