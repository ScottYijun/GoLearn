//07decision2.go
package main

import "fmt"

func main() {
	//定义局部变量
	var a int = 100
	
	//判断布尔表达式
	if (a < 20) {
		//如果条件为 true 则执行以下语句
       fmt.Printf("a 小于 20\n" )
	} else {
		//如果条件为 false 则执行以下语句
		fmt.Printf("a 不小于 20\n" );
	}
	fmt.Printf("a的值为： %d\n", a)
}

	