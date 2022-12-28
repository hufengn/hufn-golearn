package main

import "fmt"

// 数组的初始化方式
func main() {
	//第一种
	var arr1 [3]int = [3]int{3, 6, 9}
	fmt.Println(arr1)
	//第二种
	var arr2 = [3]int{1, 4, 7}
	fmt.Println(arr2)
	//第三种
	var arr3 = [...]int{4, 5, 6, 7}
	fmt.Println(arr3)
	//第四种
	var arr4 = [...]int{2: 66, 0: 33, 3: 88, 1: 99}
	fmt.Println(arr4)
}
