package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning Fine Sir!")
	case t.Hour() < 15:
		fmt.Println("Why what a fine Afternoon we are having Madam")
	default:
		fmt.Printf("Well since its later than 15:00 its evening, Actually the time is %v:%v", t.Hour(), t.Minute())
	}
}
