//04variables3.go
package main

var x, y int
var (	//这种只能出现在全局变量中，函数体内不支持
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 54, "Camel"

//这种不带声明格式的只能在函数体中出现
//g, h := 68, "Rain"

func main() {
	g, h := 68, "Rain"
	println(x, y, a, b, c, d, e, f, g, h)
}
