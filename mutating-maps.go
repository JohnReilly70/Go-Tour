package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Key"] = 42
	fmt.Println("Value: ", m["Key"])

	m["Key"] = 24
	fmt.Println("Value: ", m["Key"])

	delete(m, "Key")
	fmt.Println("Value: ", m["Key"])

	Value, isEmptyMap := m["Key"]
	fmt.Println("Value: ", Value, "Map NOT Empty?", isEmptyMap)
}
