//09functions2.go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Mahesh", "Camel")
	fmt.Println(a, b)
}

	