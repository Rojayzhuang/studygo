package main

import "fmt"

//if条件判断

func main() {
	age := 40
	if age > 18 {
		fmt.Println("大于18岁")
	} else {
		fmt.Println("<=18")
	}

	//多个条件需要判断
	if age < 35 {
		fmt.Println("大于35")
	}else if age == 40{
		fmt.Println("壮年")
	}else {
		fmt.Println("<=35")
	}

}

func ifDemo2() {
	if score := 65; score >= 90 { //添加了一个定义score的执行语句
		fmt.Println("A")		//该score仅作用于当前的if条件判断内
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}