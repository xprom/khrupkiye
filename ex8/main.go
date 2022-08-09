package main

import (
	"fmt"
	"strings"
)

func main() {
	// текст из журнала               "погода в москве 10 градусов тепла. принесите кофе"
	// нужно составить такую строку   "принесите 10" их оригинальное

	fmt.Println(
		canCreateNewString(
			"погода в москве 10 градусов градусов тепла. принесите кофе",
			"принесите 10",
		),
	)
}

func canCreateNewString(magazine string, targetString string) bool {
	list := strings.Split(magazine, " ")

	countWordsMap := make(map[string]int, len(list))
	for _, word := range list {
		if _, ok := countWordsMap[word]; ok {
			countWordsMap[word] = countWordsMap[word] + 1
		} else {
			countWordsMap[word] = 1
		}
	}

	for _, word := range strings.Split(targetString, " ") {
		count, ok := countWordsMap[word]
		if !ok {
			return false
		}

		if count == 0 {
			return false
		}

		countWordsMap[word] = countWordsMap[word] - 1
	}

	return true
}
