package main

import "fmt"

func main() {

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	//这种写法不行啊
	//mean=float32(sum/count)
	fmt.Printf("mean 的值为: %f\n", mean)

}
