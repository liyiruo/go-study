package main

import "fmt"

/**
*Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
*在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
 */
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	sum := 0
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	for _, val := range nums {
		sum += val
	}
	fmt.Printf("sum==%d\n", sum)

	for i, v := range nums {
		if i == 3 {
			fmt.Printf("i=3时，对应的value=%d\n", v)
		}
	}

	//range也可以用在map的键值对上。
	keys := map[string]string{"a": "1", "b": "2", "c": "3"}
	for k, v := range keys {
		fmt.Printf("key为%s,value 为%s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "abc" {
		fmt.Println(i, "=", c)
	}

}
