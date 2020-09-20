package main

import "fmt"

func main() {

	var numbers []int
	printSlice(numbers)
	//追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	//追加一个值
	numbers = append(numbers, 1)
	printSlice(numbers)

	//追加多个值
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
