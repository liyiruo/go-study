package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	/* 打印原始切片 */
	fmt.Println("numbers==", numbers)

	//截取前4位 0123
	fmt.Println("numbers[:4]==", numbers[:4])
	//截取前2到4位123
	fmt.Println("numbers[1:4]==", numbers[1:4])
	//截取第4位到最后
	fmt.Println("numbers[4:]==", numbers[4:])

	//创建了一个切片
	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

//
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
