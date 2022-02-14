package main

import (
	"fmt"

	"github.com/q1mi/hello" //在自己的项目中引用了别人的库
)

//函数外只能放标识符（变量、常量、函数、类型）的声明，不能防语句
//fmt.Printf("Hello")   非法

func main() {
	fmt.Printf("a")
	fmt.Println("Hello World!")
	hello.SayHi()
}
