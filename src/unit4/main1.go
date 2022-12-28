package main

import "fmt"

func main() {
	var username string
	var password string
	fmt.Println("username:")
	_, _ = fmt.Scanf("%s \n", &username)
	fmt.Printf("password: ")
	_, _ = fmt.Scanf("%s", &password)
	if username == "admin" {
		if password == "mypass" {
			fmt.Println("登录成功")
		}
	}
}
