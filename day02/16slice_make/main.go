package main

import "fmt"

//make 函数创造切片
func main() {
	slice1 := make([]int, 5, 10) //长度5，容量10,如果不设定10容量，默认和长度一样
	fmt.Printf("slice1= %v，长度：%v, 容量：%v\n", slice1, len(slice1), cap(slice1))

	slice2 := make([]int, 0, 10)
	fmt.Printf("slice2= %v，长度：%v, 容量：%v\n", slice2, len(slice2), cap(slice2))

	//切片的复制
	slice3 := []int{1, 3, 5}
	slice4 := slice3 //slice3和slice4都指向了同一个底层数组
	fmt.Println(slice4)
	slice3[0] = 100
	fmt.Println(slice3)
	fmt.Println(slice4)

	//切片的遍历
	//索引遍历
	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}
	//for range循环
	for i, v := range slice3 {
		fmt.Println(i, v)
	}
	
}
