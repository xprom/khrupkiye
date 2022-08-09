package main

import "fmt"

func main() {
	for true {
		var myName string
		var myStatus string
		var varBreak bool = false

		fmt.Println("Как тебя зовут?")
		fmt.Scan(&myName)
		fmt.Println("Привет " + myName + "!")

		fmt.Println("Как твои дела?")
		fmt.Scan(&myStatus)
		switch myStatus {
		case "Хорошо", "хорошо":
			{
				fmt.Println("Отлично!")
				varBreak = true
			}
		case "Плохо", "плохо":
			{
				fmt.Println("Печально...")
				varBreak = true
			}
		default:
			{
				fmt.Println("Не понял ответа... Давай повторим!")
			}
		}

		if varBreak == true {
			break
		}
	}

	fmt.Println("Удачного дня!")
}
