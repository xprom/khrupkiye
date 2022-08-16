package main

import (
	"fmt"
	"time"
)

func main() {
	dateTimeStart := time.Now()

	myText := "string is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not nil. Values of string type are immutable."
	fmt.Println(countingLetters(myText))

	dateTimeFinish := time.Now()
	fmt.Println("Время работы -", dateTimeFinish.Sub(dateTimeStart))
}

func countingLetters(myText string) map[string]int {
	myMap := map[string]int{}

	for a := 0; a < len(myText); a++ {
		myLitter := string([]rune(myText)[a : a+1])

		myMap[myLitter]++

		if a == len([]rune(myText))-1 {
			break
		}
	}
	return myMap
}
