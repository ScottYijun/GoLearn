package main


import (
	"fmt" 
	"strings"
)

var isHave bool	//全局变量声明

func main() {
	//对于数字类型，无需定义int及float32, float64,系统会自动识别
	var a = 45.8
	var b = 90
	fmt.Println(a, b)
	
	var isAvailable bool	//一般声明
	valid := false			//简单声明（没有写类型，这与python语言相似）
	isAvailable = true;
	//在 Go 中，布尔值的类型为 bool，值是 true 或 false，默认为 false。
	fmt.Println(isHave, isAvailable, valid)
	
	str := "这是 https://www\n.w3cschool\n.cn/go/\ngo-basic-syntax.html  的学习网址"
	fmt.Println(str)
	//去换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println(str)
	//去空格
	str = strings.Replace(str, " ", "", -1)
	fmt.Println(str)
}
	