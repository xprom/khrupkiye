package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello Gol12312312312312d!"
	fileDesc, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer func() {
		fmt.Println("запустился defer")
		fileDesc.Close()
	}()

	fileDesc.WriteString(text)

	fmt.Println("Done.")
}
