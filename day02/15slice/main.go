package main

import "fmt"

//切片
func main() {
	//切片的定义
	var slice1 []int    //定义了一个存放int类型元素的切片
	var slice2 []string //定义了一个存放string类型元素的切片

	fmt.Println(slice1, slice2)
	fmt.Println(slice1 == nil) //true,此时为空，没有开辟内存空间
	fmt.Println(slice2 == nil) //true

	//切片的初始化
	slice1 = []int{1, 2, 3, 4}
	slice2 = []string{"北京", "上海"}
	fmt.Println(slice1, slice2)
	fmt.Println(slice1 == nil) //false
	fmt.Println(slice2 == nil) //false

	//长度和容量
	fmt.Printf("len(slice1):%d \t cap(slice1):%d\n", len(slice1), cap(slice1))
	fmt.Printf("len(slice2):%d \t cap(slice2):%d\n", len(slice2), cap(slice2))

	//由数组得到切片
	array1 := [...]int{1, 3, 5, 7, 8, 6, 9, 10, 11}
	slice3 := array1[0:4] //[1,3,5,7]	基于数组的切割，左包含右不包含 [0,4)
	fmt.Println(slice3)
	slice4 := array1[2:5]
	fmt.Println(slice4)
	slice5 := array1[:]  // =>[0:len(array1)]
	slice6 := array1[:3] // =>[0:4],从0切到4
	slice7 := array1[2:] // =>[2:len(array1)]
	fmt.Println(slice5, slice6, slice7)
	//切片的容量是指：从所切的第一个元素到数组最后一个元素的数量
	fmt.Printf("len(slice6):%d \t cap(slice6):%d\n", len(slice6), cap(slice6))
	fmt.Printf("len(slice7):%d \t cap(slice6):%d\n", len(slice7), cap(slice7))
	//从切片中再切片
	slice8 := slice7[1:3]
	fmt.Println(slice8)

}
