package main

import "fmt"

//指针 pointer

func main() {
	// 两个符号
	// 1. &  ：取地址
	nums := 18
	fmt.Println(&nums)
	p := &nums
	fmt.Println(p)
	fmt.Printf("%T\n", p) // *int  *表示指针，该输出为int类型的指针
	// 2. *  ：根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) //int类型
}