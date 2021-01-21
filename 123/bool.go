package main

import "fmt"

func main()  {
   str := "steven"
   fmt.Printf("%T , %v \n", str, str)
   var a rune = '一'
   fmt.Printf("%T , %v \n", a, a)
   var b byte = 'b'
   fmt.Printf("%T , %v \n", b, b)
   var c int32 =98
   fmt.Printf("%T , %v", c, c)
}