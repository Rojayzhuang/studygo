package main

import "fmt"

//array 数组
//数组是存放元素的容器
//必须指定存放元素的类型和容量（即长度）
//数组的长度也是数组类型的一部分
//如下面array1和array2，这两个虽然都是存放bool类型元素的数组，但由于长度不同，也是不同类型的数组，因此也不能相互比较
func main() {

	var array1 [3]bool = [3]bool{true, false, false}
	var array2 [4]bool

	fmt.Printf("array1 = %T\n  array2 = %T\n", array1, array2)

	//数组的初始化
	//如果不初始化，默认数组内元素都为零值
	//布尔值就是false；浮点型、整型都是0；字符串：""
	fmt.Println(array1)
	fmt.Println(array2)

	//初始化方式1
	fmt.Println("初始化方式1：")
	var a1 [2]bool
	a1 = [2]bool{true, true}
	a2 := [3]int{1, 3, 5}
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)

	//初始化方式2
	//根据初始值自行判断长度，使用 ...
	fmt.Println("初始化方式2:")
	a3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("a3:", a3)

	//初始化方式3
	//根据索引进行初始化 索引:数值
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println("a4:", a4)

	//数组的遍历
	//根据索引遍历
	citys := [...]string{"北京", "上海", "深圳", "广州", "南京"}
	for i := 0; i < len(citys); i++ {
		fmt.Printf("city%d: %v\n", i, citys[i])
	}
	//for range遍历
	for i, chr := range citys {
		fmt.Printf("city%v: %v\n", i, chr)
	}

	//多维数组
	//[1 2] [3 4] [5 6]
	var mutiArray [3][2]int

	mutiArray = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	//fmt.Println(mutiArray)

	//多维数组的遍历
	for _, v1 := range mutiArray {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型而不是引用类型
	b1 := [3]int{0, 1, 2} //b1 = [0 1 2]
	b2 := b1              //b2 = [0 1 2] 相当于复制操作，并非引用
	fmt.Println(b1, b2)

	b2[0] = 100         //b2 = [100 1 2]
	fmt.Println(b1, b2) //b1 = [0 1 2]

	//习题：求数组[1,3,5,7,8]所有元素的和
	workArray := [...]int{1, 3, 5, 7, 8}
	sumArray := 0
	for _, v := range workArray {
		sumArray += v
	}
	fmt.Println(sumArray)

	//习题：找出数组中和为指定值的两个元素的下标，
	//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	sumArrayValue(workArray[:], 8)
}

//习题：找出数组中和为指定值的两个元素的下标，
//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func sumArrayValue(array []int, sumValue int) {
	for i := 0; i < len(array); i++ {
		for c := len(array) - 1; c >= 0; c-- {
			if array[i]+array[c] == sumValue {
				fmt.Printf("(%d,%d)", i, c)
			} else if i == c {
				break
			}
		}
	}
}
