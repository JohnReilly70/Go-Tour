package main

import "fmt"

func swap (x,y string) (string,string){
  return y,x
}

func main(){
  a, b :=swap("Hi","Bye")
  fmt.Println(a,b)
}
