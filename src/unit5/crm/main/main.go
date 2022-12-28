package main

import "fmt"
import "hufn-golearn/src/unit5/crm/dbutils"

// import(
// "fmt"
// "hufn-golearn/src/unit5/crm/dbutils"
// )
// 另一种导入包的方法
func main() {
	fmt.Println("你好，这是main函数的执行")
	dbutils.GetConn()
}
