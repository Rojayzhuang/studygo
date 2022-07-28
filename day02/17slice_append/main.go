package main

import "fmt"

//切片的 append() 为切片追加元素

func main() {
	slice1 := []string{"北京", "上海", "深圳"}
	//slice1[3] = "广州" //错误写法，编译错误，索引越界
	fmt.Printf("slice1:%v , len(slice1):%d, cap(slice1):%d\n",
		slice1, len(slice1), cap(slice1))

	//调用append函数必须用原来的切片变量接收返回值
	//append追加元素后，原来的底层数组放不下时，go语言底层就会把底层数组换一个
	//必须用变量接收append的返回值
	slice2 := append(slice1, "临沂")
	fmt.Printf("slice2:%v , len(slice2):%d, cap(slice2):%d\n",
		slice2, len(slice2), cap(slice2))
	slice1 = append(slice1, "广州", "南京", "南京", "南京")
	fmt.Printf("slice1:%v , len(slice1):%d, cap(slice1):%d\n",
		slice1, len(slice1), cap(slice1))

	//append一个切片
	sliceAppend := []string{"杭州", "武汉"}
	//使用...   ...表示把前面的切片拆开
	slice1 = append(slice1, sliceAppend...)
	fmt.Printf("slice1:%v , len(slice1):%d, cap(slice1):%d\n",
		slice1, len(slice1), cap(slice1))

}
