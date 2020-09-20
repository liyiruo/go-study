package main

import "fmt"

func main() {
	var a int = 20
	var b *int
	b = &a
	fmt.Printf("a变量的地址是：%x\n", b)
	fmt.Printf("a的值：%d\n", *b)

	var c *int
	fmt.Printf("b的值为：%x\n", c)
	if c == nil {
		fmt.Print("c为空")
	}
	if c != nil {
		fmt.Print("c不为空")
	}
}
