package main

import (
	"fmt"
	"sort"
)

//切片练习题
func main() {
	var slice1 = make([]int, 5, 10) //创建切片，长度为5，容量为10
	fmt.Println("slice1: ", slice1) //[0 0 0 0 0]
	for i := 0; i < 10; i++ { 
		slice1 = append(slice1, i)
	}
	fmt.Println("slice1添加后: ", slice1)
	fmt.Println(cap(slice1))

	var	array1 = [...]int{3,7,8,6,1}
	fmt.Println(array1)
	sort.Ints(array1[:]) //对切片进行排序
	fmt.Println(array1)
}