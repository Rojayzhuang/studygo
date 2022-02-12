package main

import (
	"fmt"

	"github.com/q1mi/hello" //在自己的项目中引用了别人的库
)

func main() {
	fmt.Println("123")
	hello.SayHi()
}
