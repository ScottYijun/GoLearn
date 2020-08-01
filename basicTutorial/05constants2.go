//05constants2.go
package main

import "unsafe"

const (
	a = "Camel"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)
}
	