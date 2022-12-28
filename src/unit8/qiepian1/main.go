package main

import "fmt"

func main() {
	//定义数组
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	slice := intarr[1:3]
	//输出数组：
	fmt.Println("intarr", intarr)
	//输出切片
	fmt.Println("slice", slice)
	//切片元素个数
	fmt.Println("slice的元素个数：", len(slice))
	//获取切片的容量：容量可以动态变化
	fmt.Println("slice的容量：", cap(slice))
}
