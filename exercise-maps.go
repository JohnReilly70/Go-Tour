package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	Word_Map := make(map[string]int)
	Word_List := strings.Fields(s)

	for _, word := range Word_List {
		fmt.Println(word)
		_, Key_Used := Word_Map[word]
		fmt.Println(Key_Used)
		if Key_Used {
			Word_Map[word]++
		} else {
			Word_Map[word] = 1
		}

	}
	fmt.Println(Word_Map)
	return Word_Map
}

func main() {
	WordCount("Hello Hello HELLO BOO")
}
