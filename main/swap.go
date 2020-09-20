package main

import "fmt"

//定义一个 函数

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	m, n := swap("111", "222")
	fmt.Println(m, n)
}
