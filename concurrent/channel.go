package main

import "fmt"

func main() {
	//创建一个数组
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	//计算出的值为-9+4+0=-5
	go sum(s[:len(s)/2], c)
	//计算出的值为7+2+8=17
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

//创建一个方法
func sum(s []int, c chan int) {
	sum := 0
	//获取s里的值的和
	for _, v := range s {
		sum += v
	}
	//把sum值付给c
	c <- sum

}
