package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func sub(x, y int) int {
  return x - y
}

func mult(x,y float32) float32 {
  return x*y
}

func main() {

  fmt.Println(add(42,2))
  fmt.Println(sub(42,2))
  fmt.Println(mult(42,2))
}
