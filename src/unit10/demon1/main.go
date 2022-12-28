package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	var t1 Teacher
	t1.Name = "马士兵"
	t1.Age = 45
	t1.School = "清华大学"
	fmt.Println(t1)
}
