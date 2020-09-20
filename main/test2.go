package main

import "fmt"

/*声明全局变量*/
var g int = 200

func main() {
	//声明局部变量
	var a, b int
	a = 100
	b = 200
	g = a + b
	fmt.Printf("g= %d", g)
}
