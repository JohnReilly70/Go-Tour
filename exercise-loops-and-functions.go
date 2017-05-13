package main

import (
	"fmt"
)

func Sqrt_1_eq(answer, x float64) float64 {
	return answer*answer - x
}
func Sqrt_2_eq(answer float64) float64 {
	return 2 * answer
}

func Sqrt(x float64) float64 {
	var answer float64 = x / 2
	var func1, func2 float64 = 0, 0
	for i := 0; i < 10; i++ {
		func1 = Sqrt_1_eq(answer, x)
		//fmt.Println(func1)
		func2 = Sqrt_2_eq(answer)
		//fmt.Println(func2)
		answer = answer - (func1 / func2)
		//fmt.Println(answer)
	}
	return answer
}

func main() {
	fmt.Printf("Square root of 9: %v\n", Sqrt(9))
	fmt.Printf("Square root of 12: %v\n", Sqrt(12))
}
