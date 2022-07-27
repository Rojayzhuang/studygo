package main

import "fmt"

//make 函数创造切片
func main() {
	slice1 := make([]int, 5, 10) //长度5，容量10,如果不设定10容量，默认和长度一样
	fmt.Printf("slice1= %v，长度：%v, 容量：%v\n", slice1, len(slice1), cap(slice1))

	slice2 := make([]int, 0, 10)
	fmt.Printf("slice2= %v，长度：%v, 容量：%v\n", slice2, len(slice2), cap(slice2))
}
