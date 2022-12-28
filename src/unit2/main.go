package main

import "fmt"

func main() {
	var score int = 96
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("B")
	} else if score < 60 {
		fmt.Println("B")
	}
}
