package main

import "fmt"

// 函数：
// 对特定的功能进行提醒，形成一个代码片段，这个代码片段就是我们所说的函数
// 函数的作用：1. 提高代码的复用性   2. 函数和函数是并列的关系，所以我们定义的函数不能写到main 函数中。
// 函数语法：
//
//	func 函数名 （形参列表）（返回值类型列表）{
//		执行语句
//		return + 返回值列表
//	}

//函数名：
/*遵循标识符命名规范：见名知意 addNum,驼峰命名addNum
首字母不能是数字
首字母大写该函数可以被本包文件和其他包文件使用（类似public）
首字母小写只能被本报文件使用，其他包问渐渐不能使用（类似private）*/
func cal(num1 int, num2 int) int {
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}
func main() {
	//var num1 int = 10
	//var num2 int = 20
	//var sum int = 0
	sum := cal(10, 20)
	fmt.Println(sum)
	//sum += num1
	//sum += num2
	//fmt.Println(sum)

}
