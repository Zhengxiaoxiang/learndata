package main

import "fmt"

func main()  {
	//声明一个变量并且初始化
	var a = "zxx"
	fmt.Print(a)

	//没有初始化就为零值
	var b int
	fmt.Print(b)

	//bool 零值为false
	var c  bool
	fmt.Print(c)
}