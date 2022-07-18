package main

import (
	"fmt"
	"strings"
)

//字符串

func main() {
	path := "\\Users\\rojayzhuang\\Downloads"

	//引号的转义
	path2 := "\"\\Users\\rojayzhuang\\Downloads\""
	fmt.Printf("%v\n",path)
	fmt.Printf("%v\n", path2)
	//也可以使用#v的方式输出值，这样自带引号
	fmt.Printf("%#v\n", path)
	//单行字符串
	s1 := "this is a string"
	fmt.Printf("%v\n", s1)
	//多行字符串
	s2 := `举头望明月
低头思故乡`
	fmt.Println(s2)	
	path3 := `\Users\rojayzhuang\Downloads`
	fmt.Println("path3:",path3)

	//字符串相关操作

	//拼接字符串
	s3 := "i am "
	s4 := "rojayzhuang"
	fmt.Println(len(s3))
	fmt.Println(s3+s4)

	ss := fmt.Sprintf("%v%v",s3,s4)
	fmt.Println(ss)

	//分割字符串

	split1 := strings.Split(path3, "\\")
	fmt.Printf(split1[1]+split1[3])
	fmt.Println(len(split1))

	//join拼接
	fmt.Println(strings.Join(split1,"+"))

	s5 := "abcdebcdebcde"
	result := strings.LastIndex(s5,"bc")
	fmt.Println(result)

	str2:= "GFG is the Best"
	fmt.Println(strings.IndexAny(str2,"Be"))
	
}
