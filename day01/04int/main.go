package main

import "fmt"

//整形

func main() {
	//十进制
	var i1 = 101
	fmt.Printf("%d\n", i1) //%d 来表示一个十进制的数
	fmt.Printf("%b\n", i1) //%b表示二进制（把十进制转换成二进制）
	fmt.Printf("%o\n", i1) //%o表示八进制 （把十进制转换成八进制）
	fmt.Printf("%x\n", i1) //%x表示十六进制 （把十进制转换成十六进制）
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	//查看变量的类型 %T
	fmt.Printf("%T\n", i3)

	//声明int8类型的变量
	i4 := int8(9)  //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)

}
