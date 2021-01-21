package main

import "fmt"

func main()  {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a //a copy of a s assigned
	b[0] = "shenzhen"
	fmt.Println("a :", a)
	fmt.Println("b : ", b)
}