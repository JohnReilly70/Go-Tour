package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("My random number is", rand.Intn(200))
}
