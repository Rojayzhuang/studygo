package main

import (
	"fmt"
	"unicode"
)

func main() {

//traveralString()
fixString("123abc")
countChineseCharactersNumber("123abc我爱中国")

}


func traveralString(){
	str := "hello庄"
	for i := 0; i <len(str); i++ {
		fmt.Printf("%v(%c)", str[i], str[i])
	}
	fmt.Println()
	for _, r := range str {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}
//修改字符串
func fixString(str string) {
	fmt.Println(str)
	//字符串强制转换为一个rune切片
	fixStr := []rune(str)
	fixStr[1] = '改'
	fmt.Println(string(fixStr))
	fmt.Printf("%T\n", str)
	fmt.Printf("%T\n", fixStr)
}

//统计汉字数量
func countChineseCharactersNumber(str string) {
	fmt.Printf("需要统计汉字字数的字符串：%v\n",str)
	
	runeStr := []rune(str)
	countChinese := 0;
	for _, chr := range runeStr {
		//unicode.Han go语言自带的判断中文
		if unicode.Is(unicode.Han, chr) {
			countChinese++
		}
	}
	
	fmt.Printf("%v\n",countChinese)
		
	}
