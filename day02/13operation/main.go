package main

import (
	"fmt"
)

func main() {

	//取反运算符，原来为真就为假
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	//位运算符，针对二进制计算
	//5的二进制：101
	//2的二进制：10
	fmt.Println(5 ^ 2)

	//有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
	// 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
	// 相同，为 0
	nums := [...]int{1, 2, 3, 3, 2, 1, 6, 9, 8, 8, 6}

	// 	任何数和 00 做异或运算，结果仍然是原来的数，即 a ^ 0=aa⊕0=a
	//  任何数和其自身做异或运算，结果是 00，即 a ^ a=0a⊕a=0
	//  异或运算满足交换律和结合律，即：
	// a ^ b ^ a=b ^ a ^ a=b ^ (a ^ a)=b ^ 0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。

	//原理为异或
	fmt.Println(nums[0] ^ nums[1] ^ nums[2] ^ nums[3] ^ nums[4] ^ nums[5] ^ nums[6] ^ nums[7] ^ nums[8] ^ nums[9] ^ nums[10])

	//使用循环
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		//n ^= nums[i]
		n = n ^ nums[i]
		//1 = 1 ^ nums[1]
	}

	fmt.Println(n)

}
