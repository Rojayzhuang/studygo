package main

import "fmt"

//常量学习
//定义了常量后不能修改
//在程序运行期间不能修改

const pi = 3.141592653589

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

//如果批量声明常量时，某一行没有赋值，那么默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2        //1
	_         //2,匿名变量
	b3        //3
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3        //100
	c4 = iota //3  因为在同一个const中，所以共用一个计数
	c5        //4
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1:1, d2:2

	d3, d4 = iota + 1, iota + 2 //d3:2, d4:3
)

//定义数量级
const (
	_  = iota //
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)

	fmt.Println()

	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)

	fmt.Println()

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)

	fmt.Println()

	fmt.Println("c1: ", c1)
	fmt.Println("c2: ", c2)
	fmt.Println("c3: ", c3)
	fmt.Println("c4: ", c4)
	fmt.Println("c5: ", c5)

	fmt.Println()

	fmt.Println("d1: ", d1)
	fmt.Println("d2: ", d2)
	fmt.Println("d3: ", d3)
	fmt.Println("d4: ", d4)
}
