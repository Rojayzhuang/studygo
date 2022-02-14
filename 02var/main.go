package main

import "fmt"

//声明变量

var name string
var age int
var isOk bool

//批量声明
// var (
// 	name string
// 	age  int
// 	isOk bool
// )

func main() {
	name = "李想"
	age = 18
	isOk = true
	//var bianliang string
	//go语言中非全局变量声明了必须使用
	
	fmt.Print(isOk)
	fmt.Printf("name:%s", name) //%s 占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)


	//声明变量同时赋值
	var s1 string = "黎明"
	fmt.Println(s1)

	//类型推导，根据值判断该变量是什么类型
	var s2 = "2022"
	fmt.Println(s2)

	//短变量声明，仅能在函数中使用
	s3 := "简短变量声明"
	fmt.Println(s3)



}
