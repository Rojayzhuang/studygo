package main

import "fmt"

//for循环

func main() {
	//基本格式
	// for i:=0; i<10; i++ {
	// 	fmt.Println(i)
	// }

	//无限循环
	// for {
	// 	fmt.Println("123")
	// }

	//for range 循环
	str := "hello world你好世界"

	for i,v := range str{
		fmt.Printf("%d %c\n", i, v)
	}

	//9*9乘法表
	for a := 1 ; a < 10; a++ {
		for b := 1; b <= a ; b++ {
		fmt.Printf("%v * %v = %v\t" , a ,b, (a*b)) //使用\t制表符，自动输出tab
	}
	fmt.Println()

	//break 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("Break 跳出，i=%d",i)
			break
		}
		fmt.Println(i)
	}


	//continue 当i=5时，跳过此次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
	

}