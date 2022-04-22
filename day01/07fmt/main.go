package main

import "fmt"

//fmt占位符总结

func main() {
	var n = 100
	s1 := "字符串123"
	fmt.Printf("%T\n", n)  //查看类型
	fmt.Printf("%#v\n", n) //查看值
	fmt.Printf("%b\n", n)  //查看二进制
	fmt.Printf("%o\n", n)  //查看八进制
	fmt.Printf("%d\n", n)  //查看十进制
	fmt.Printf("%x\n", n)  //查看十六进制

	fmt.Printf("%s\n", s1)        //查看字符串
	fmt.Printf("字符串的值：%#v\n", s1) //查看值,#可以对输出的值进行描述，例如字符串会给加上双引号
}
