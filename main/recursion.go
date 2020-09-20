package main

import "fmt"

//写一个递归吧
//给一个数字 计算这个数字的递归值
func main() {
	num := 5
	s := recursion(num)
	fmt.Printf("num的阶乘为：%d", s)
}

//传入的是a是什么类型的，返回的是什么类型的
func recursion(a int) int {
	if a <= 0 {
		return 0
	} else if a == 1 {
		return 1
	} else {
		return a * recursion(a-1)
	}
}

//另一种写法
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
