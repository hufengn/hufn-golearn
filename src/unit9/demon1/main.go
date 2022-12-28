package main

import "fmt"

func main() {
	var a map[int]string
	// 只声明map内存是没有分配空间
	//方式1
	a = make(map[int]string, 10)
	//map可以存放10个键值对
	a[20095452] = "张三"
	a[20095451] = "李四"
	a[20097291] = "王五"
	fmt.Println(a)
	//方式2
	b := make(map[int]string)
	b[20095452] = "张三"
	b[20095451] = "李四"
	fmt.Println(b)
	//方式3
	c := map[int]string{
		20095452: "张三",
		20095451: "李四",
	}
	c[20095452] = "张三"
	fmt.Println(c)
	//map 的特点：
	// map集合在使用前一定要make
	// map的key-value是无序的
	// key是不可以重复的，如遇到重复，后一个value 会替换前一个value
	// value  可以重复
	//键值对数量可以省略
}
