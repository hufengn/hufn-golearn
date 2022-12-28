package dbutils

import "fmt"

func GetConn() { //首字母大写，其他包才能调用
	fmt.Println("执行了dbutils包下的getConn函数")
}
