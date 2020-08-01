//05constants.go
package main

import "fmt"

func main() {
	const LENGTH int = 30
	const WIDTH	 int = 90
	var area int
	const a, b, c = 1, false, "Camel"	//多重赋值
	area = LENGTH * WIDTH
	fmt.Printf("面积为： %d\n", area)
	
	println(a, b, c)
}

	