package main

import (
	"fmt"
	"math"
)

//浮点数
func main() {
	fmt.Println(math.MaxFloat32)// float32最大值

	f1 := 1.23456
	fmt.Printf("f1=1.23456的数据类型为:%T\n",f1) // 默认go语言中的小数都是float64类型
	f2 := float32(1.23456) //定义了f2的数据类型
	fmt.Printf("显示声明float32类型:%T\n",f2) 
	//f1 = f2 //类型不同，不能相互赋值  float32不能直接赋值给float64
}