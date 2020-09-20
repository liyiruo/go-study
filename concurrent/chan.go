package main

import "fmt"

//通道缓冲区
func main() {

	//创建一个缓冲区，缓冲区大小为2，缓冲区的类型为int
	ints := make(chan int, 2)
	//往缓冲区里放数据
	ints <- 5
	ints <- 10
	//从缓冲区里取数据
	a := <-ints
	b := <-ints
	fmt.Println(a)
	fmt.Println(b)
}
