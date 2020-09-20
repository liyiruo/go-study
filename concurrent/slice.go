package main

import "fmt"

func main() {
	//初始化
	//var numbers=make([]int,3,5)
	//未初始化
	var numbers []int
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

//创建一个方法 获取某个切面的长度，容量 和切面的内容
func printSlice(x []int) {
	fmt.Printf("len=%d,cap=%d,slice=%v \n", len(x), cap(x), x)
}
