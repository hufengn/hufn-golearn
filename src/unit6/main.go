package main

import "fmt"

// 排错机制
// 数组遍历
func main() {
	var scores [5]int
	//
	for i := 0; i < len(scores); i++ {
		fmt.Println("请录入第%d个学生的成绩", i+1)
		fmt.Scanln(&scores[i])
	}
	//数组的遍历
	for key, value := range scores {
		fmt.Printf("第%d个学生的成绩为： %d\n", key, value)
	}
}
