package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func sub(x int, y int) int {
  return x - y
}

func main() {

  fmt.Println(add(42,2))
  fmt.Println(sub(42,2))
}
