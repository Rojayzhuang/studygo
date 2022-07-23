package main

import "fmt"

//switch case 和goto

func main() {
	//switch 简化大量判断
	switch n := 0; {
	case n%2 == 1:
		fmt.Println("奇数")
	default:
		fmt.Println("偶数")
	}

	switch n := 6; n {
	case 1, 3, 5, 7:
		fmt.Println("奇数")
	case 2, 4, 6:
		fmt.Println("偶数")
	}

	//goto 跳到指定标签
	//未使用goto+lable跳出多层循环
	var flag = false //定义一个布尔值，用作标签判断
	for i := 0; i < 10; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if j == 'C' {
				flag = true //要跳出外层循环
				break       //仅能跳出内层循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			fmt.Println("over")
			break //跳出外层循环
		}
	}

	//上述使用goto + lable后
	for i := 0; i < 10; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if j == 'C' {
				goto breakLabel //跳到指定标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	breakLabel:  //标签
		fmt.Println("over")

}
