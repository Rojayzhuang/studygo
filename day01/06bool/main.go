package main

import "fmt"

//bool

func main(){
	b1 := true
	var b2 bool
	fmt.Printf("b1的数据类型：%T\n", b1)
	fmt.Printf("b2的数据类型：%T\n", b2)
	fmt.Printf("%T value:%v\n",b2, b2)  //%v打印变量的值
}