package main

import "fmt"

//copy

func main() {
	slice1 := []int{1, 3, 5}
	slice2 := slice1
	var slice3 []int                    //这样声明为nil，没有空间，不能copy
	copy(slice3, slice1)                //copy(需要copy的数组、切片 , copy的数组、切片)
	fmt.Println(slice1, slice2, slice3) //[1 3 5] [1 3 5] []
	var slice4 = make([]int, 3, 3) //这样可以copy
	copy(slice4, slice1) //copy
	fmt.Println(slice1, slice2, slice4) //[1 3 5] [1 3 5] [1 3 5]
	slice1[0] = 100
	fmt.Println(slice1, slice2, slice4) //[100 3 5] [100 3 5] [1 3 5]

	//从切片中删除元素
	//go语言中没有删除切片元素的专用方法，可以使用切片本身的特性来删除元素
	sliceDelete := []int{30, 31, 32, 33, 34, 35, 36, 37, 38, 39}
	//删除索引为2的元素
	sliceDelete = append(sliceDelete[:2], sliceDelete[3:]...)
	fmt.Println(sliceDelete)

	//奖slice1索引为1的元素3删掉
	//修改了底层数组
	//...表示拆开
	append1 := append(slice1[:1], slice1[2:]...)
	fmt.Println(append1)
	//切片不存值，底层数组存值
	fmt.Println(cap(slice1), cap(append1), slice1)

}
