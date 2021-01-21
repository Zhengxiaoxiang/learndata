package main

import "fmt"

func main()  {
	/* 创建切片 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)
	fmt.Println("number ==", numbers)
	/*打印子切片索引0(包含）到索引4（不包含）*/
	fmt.Println("numnbers[1:4] ==", numbers[1:4])
	/*默认下限为0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/*默认上限len（s）*/
	fmt.Println("numbers[4:]", numbers[4:])
	numbers2 := numbers[:2]
	printSlice(numbers2)
	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}